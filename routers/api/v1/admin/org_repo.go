// Copyright 2016 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	"github.com/gogits/gogs/models"
	"github.com/gogits/gogs/modules/context"
)

func GetRepositoryByParams(ctx *context.APIContext) *models.Repository {
	repo, err := models.GetRepositoryByName(ctx.Org.Team.OrgID, ctx.Params(":reponame"))
	if err != nil {
		if models.IsErrRepoNotExist(err) {
			ctx.Status(404)
		} else {
			ctx.Error(500, "GetRepositoryByName", err)
		}
		return nil
	}
	return repo
}

func AddTeamRepository(ctx *context.APIContext) {
	repo := GetRepositoryByParams(ctx)
	if ctx.Written() {
		return
	}
	if err := ctx.Org.Team.AddRepository(repo); err != nil {
		ctx.Error(500, "AddRepository", err)
		return
	}

	ctx.Status(204)
}

func RemoveTeamRepository(ctx *context.APIContext) {
	repo := GetRepositoryByParams(ctx)
	if ctx.Written() {
		return
	}
	if err := ctx.Org.Team.RemoveRepository(repo.ID); err != nil {
		ctx.Error(500, "RemoveRepository", err)
		return
	}

	ctx.Status(204)
}
