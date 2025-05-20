// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"

	"github.com/minherz/bnb-demo/frontend/utils"
)

type Config struct {
	Host              string
	InstanceId        string
	Location          string
	Port              string
	ProjectId         string
	RevisionName      string
	ServiceName       string
	StaticPath        string
	CatalogServiceURI string
	VideoServiceURI   string
}

const (
	DefaultPortNumber = "8080"
)

func NewConfig(ctx context.Context, getenv func(string) string) *Config {
	utils.Setup(getenv)
	catalogService := utils.FormatServiceName(ctx, utils.GetEnv("CATALOG_SERVICE", ""))
	if catalogService == "" {
		catalogService = utils.FormatServiceName(ctx, "bnb-catalog")
	}
	return &Config{
		Host:              "",
		Location:          utils.StringOnly(utils.Region(ctx)),
		Port:              utils.GetStringParam("port", "PORT", "8080"),
		ProjectId:         utils.StringOnly(utils.ProjectID(ctx)),
		RevisionName:      utils.RevisionName(),
		ServiceName:       utils.ServiceName(),
		StaticPath:        utils.GetStringParam("static", "STATIC_PATH", "./_static/"),
		CatalogServiceURI: utils.FormatServiceName(ctx, utils.GetEnv("CATALOG_SERVICE", ""), defaultCatalogService),
		VideoServiceURI:   utils.FormatServiceName(ctx, utils.GetEnv("VIDEO_SERVICE", ""), defaultCatalogService),
	}
}
