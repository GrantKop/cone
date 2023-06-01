package main

import (
	"context"
	"fmt"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
	"github.com/spf13/cobra"
)

func getCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Create an access request for an entitlement by slug",
		RunE:  runGet,
	}

	addWaitFlag(cmd)
	addAppIdFlag(cmd)
	addEntitlementIdFlag(cmd)

	return cmd
}

func runGet(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	alias := ""

	v, err := getSubViperForProfile(cmd)
	if err != nil {
		return err
	}

	clientId, clientSecret, err := getCredentials(v)
	if err != nil {
		return err
	}

	entitlementId := v.GetString(entitlementIdFlag)
	appId := v.GetString(appIdFlag)

	if len(args) == 1 {
		alias = args[0]
	}

	if alias == "" && (appId == "" || entitlementId == "") {
		return fmt.Errorf("must provide either an alias or an entitlement id and app id")
	}

	if alias != "" && (appId != "" || entitlementId != "") {
		return fmt.Errorf("cannot provide both an alias and an entitlement id and app id")
	}

	c, err := client.New(ctx, clientId, clientSecret, client.WithDebug(v.GetBool("debug")))
	if err != nil {
		return err
	}

	if alias != "" {
		entitlement, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{EntitlementAlias: alias})
		if err != nil {
			return err
		}
		if len(entitlement.List) == 0 {
			return fmt.Errorf("no entitlement found with alias %s", alias)
		}
		if len(entitlement.List) > 1 {
			// TODO: this should show a list and prompt for input.
			return fmt.Errorf("multiple entitlements found with alias %s", alias)
		}
		entitlementId = client.StringFromPtr(entitlement.List[0].AppEntitlement.Id)
		appId = client.StringFromPtr(entitlement.List[0].AppEntitlement.AppId)
	}

	resp, err := c.WhoAmI(ctx)
	if err != nil {
		return err
	}

	accessRequest, err := c.CreateGrantTask(ctx, appId, entitlementId, client.StringFromPtr(resp.UserId))
	if err != nil {
		return err
	}

	taskResp := C1ApiTaskV1Task(*accessRequest.TaskView.Task)

	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &taskResp)
	if err != nil {
		return err
	}

	return nil
}

type C1ApiTaskV1Task c1api.C1ApiTaskV1Task

func (r *C1ApiTaskV1Task) Header() []string {
	return []string{"NumericId", "Id", "DisplayName", "Description", "State", "Processing"}
}
func (r *C1ApiTaskV1Task) Rows() [][]string {
	return [][]string{{
		client.StringFromPtr(r.NumericId),
		client.StringFromPtr(r.Id),
		client.StringFromPtr(r.DisplayName),
		client.StringFromPtr(r.Description),
		client.StringFromPtr(r.State),
		client.StringFromPtr(r.Processing),
	}}
}
