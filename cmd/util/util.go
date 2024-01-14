package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/spf13/viper"
	"github.com/tidwall/pretty"
)

type WorkflowDefSummary struct {
	OwnerApp      string `json:"ownerApp,omitempty"`
	CreateTime    int64  `json:"createTime,omitempty"`
	UpdateTime    int64  `json:"updateTime,omitempty"`
	CreatedBy     string `json:"createdBy,omitempty"`
	UpdatedBy     string `json:"updatedBy,omitempty"`
	Name          string `json:"name"`
	Description   string `json:"description,omitempty"`
	Version       int32  `json:"version,omitempty"`
	SchemaVersion int32  `json:"schemaVersion,omitempty"`
	OwnerEmail    string `json:"ownerEmail,omitempty"`
}

func SummarizeWorkflowDef(wfDef *model.WorkflowDef) (WorkflowDefSummary, error) {
	if wfDef == nil {
		return WorkflowDefSummary{}, errors.New("Input WorkflowDef is nil")
	}

	// Create a summarized version
	summary := WorkflowDefSummary{
		OwnerApp:      wfDef.OwnerApp,
		CreateTime:    wfDef.CreateTime,
		UpdateTime:    wfDef.UpdateTime,
		CreatedBy:     wfDef.CreatedBy,
		UpdatedBy:     wfDef.UpdatedBy,
		Name:          wfDef.Name,
		Description:   wfDef.Description,
		Version:       wfDef.Version,
		SchemaVersion: wfDef.SchemaVersion,
		OwnerEmail:    wfDef.OwnerEmail,
	}

	return summary, nil
}

func PrintJSON(data interface{}) {
	// Marshal the entire slice into JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	// Beautify the JSON using pretty package
	prettified := pretty.Pretty(jsonData)

	// Print the JSON array
	fmt.Println(string(prettified))
}

var Authsettings = func() *settings.AuthenticationSettings {
	if viper.IsSet("key") && viper.IsSet("secret") {
		return settings.NewAuthenticationSettings(viper.GetString("key"), viper.GetString("secret"))
	}
	return nil
}()
