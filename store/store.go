package store

import (
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/model/forge"
	"github.com/kleister/kleister-api/model/minecraft"
)

//go:generate mockery -all -case=underscore

// Store implements all required data-layer functions for Solder.
type Store interface {
	// GetBuilds retrieves all available builds from the database.
	GetBuilds(int) (*model.Builds, error)

	// CreateBuild creates a new build.
	CreateBuild(int, *model.Build) error

	// UpdateBuild updates a build.
	UpdateBuild(int, *model.Build) error

	// DeleteBuild deletes a build.
	DeleteBuild(int, *model.Build) error

	// GetBuild retrieves a specific build from the database.
	GetBuild(int, string) (*model.Build, *gorm.DB)

	// GetBuildVersions retrieves versions for a build.
	GetBuildVersions(*model.BuildVersionParams) (*model.BuildVersions, error)

	// GetBuildHasVersion checks if a specific version is assigned to a build.
	GetBuildHasVersion(*model.BuildVersionParams) bool

	// CreateBuildVersion assigns a version to a specific build.
	CreateBuildVersion(*model.BuildVersionParams) error

	// DeleteBuildVersion removes a version from a specific build.
	DeleteBuildVersion(*model.BuildVersionParams) error

	// GetClients retrieves all available clients from the database.
	GetClients() (*model.Clients, error)

	// CreateClient creates a new client.
	CreateClient(*model.Client) error

	// UpdateClient updates a client.
	UpdateClient(*model.Client) error

	// DeleteClient deletes a client.
	DeleteClient(*model.Client) error

	// GetClient retrieves a specific client from the database.
	GetClient(string) (*model.Client, *gorm.DB)

	// GetClientPacks retrieves packs for a client.
	GetClientPacks(*model.ClientPackParams) (*model.ClientPacks, error)

	// GetClientHasPack checks if a specific pack is assigned to a client.
	GetClientHasPack(*model.ClientPackParams) bool

	// CreateClientPack assigns a pack to a specific client.
	CreateClientPack(*model.ClientPackParams) error

	// DeleteClientPack removes a pack from a specific client.
	DeleteClientPack(*model.ClientPackParams) error

	// GetMods retrieves all available mods from the database.
	GetMods() (*model.Mods, error)

	// CreateMod creates a new mod.
	CreateMod(*model.Mod) error

	// UpdateMod updates a mod.
	UpdateMod(*model.Mod) error

	// DeleteMod deletes a mod.
	DeleteMod(*model.Mod) error

	// GetMod retrieves a specific mod from the database.
	GetMod(string) (*model.Mod, *gorm.DB)

	// GetModUsers retrieves users for a mod.
	GetModUsers(*model.ModUserParams) (*model.UserMods, error)

	// GetModHasUser checks if a specific user is assigned to a mod.
	GetModHasUser(*model.ModUserParams) bool

	// CreateModUser assigns a user to a specific mod.
	CreateModUser(*model.ModUserParams) error

	// DeleteModUser removes a user from a specific mod.
	DeleteModUser(*model.ModUserParams) error

	// GetModTeams retrieves teams for a mod.
	GetModTeams(*model.ModTeamParams) (*model.TeamMods, error)

	// GetModHasTeam checks if a specific team is assigned to a mod.
	GetModHasTeam(*model.ModTeamParams) bool

	// CreateModTeam assigns a team to a specific mod.
	CreateModTeam(*model.ModTeamParams) error

	// DeleteModTeam removes a team from a specific mod.
	DeleteModTeam(*model.ModTeamParams) error

	// GetPacks retrieves all available packs from the database.
	GetPacks() (*model.Packs, error)

	// CreatePack creates a new pack.
	CreatePack(*model.Pack) error

	// UpdatePack updates a pack.
	UpdatePack(*model.Pack) error

	// DeletePack deletes a pack.
	DeletePack(*model.Pack) error

	// GetPack retrieves a specific pack from the database.
	GetPack(string) (*model.Pack, *gorm.DB)

	// GetPackClients retrieves clients for a pack.
	GetPackClients(*model.PackClientParams) (*model.ClientPacks, error)

	// GetPackHasClient checks if a specific client is assigned to a pack.
	GetPackHasClient(*model.PackClientParams) bool

	// CreatePackClient assigns a client to a specific pack.
	CreatePackClient(*model.PackClientParams) error

	// DeletePackClient removes a client from a specific pack.
	DeletePackClient(*model.PackClientParams) error

	// GetPackUsers retrieves users for a pack.
	GetPackUsers(*model.PackUserParams) (*model.UserPacks, error)

	// GetPackHasUser checks if a specific user is assigned to a pack.
	GetPackHasUser(*model.PackUserParams) bool

	// CreatePackUser assigns a user to a specific pack.
	CreatePackUser(*model.PackUserParams) error

	// DeletePackUser removes a user from a specific pack.
	DeletePackUser(*model.PackUserParams) error

	// GetPackTeams retrieves teams for a pack.
	GetPackTeams(*model.PackTeamParams) (*model.TeamPacks, error)

	// GetPackHasTeam checks if a specific team is assigned to a pack.
	GetPackHasTeam(*model.PackTeamParams) bool

	// CreatePackTeam assigns a team to a specific pack.
	CreatePackTeam(*model.PackTeamParams) error

	// DeletePackTeam removes a team from a specific pack.
	DeletePackTeam(*model.PackTeamParams) error

	// GetUsers retrieves all available users from the database.
	GetUsers() (*model.Users, error)

	// CreateUser creates a new user.
	CreateUser(*model.User) error

	// UpdateUser updates a user.
	UpdateUser(*model.User) error

	// DeleteUser deletes a user.
	DeleteUser(*model.User) error

	// GetUser retrieves a specific user from the database.
	GetUser(string) (*model.User, *gorm.DB)

	// GetUserMods retrieves mods for a user.
	GetUserMods(*model.UserModParams) (*model.UserMods, error)

	// GetUserHasMod checks if a specific mod is assigned to a user.
	GetUserHasMod(*model.UserModParams) bool

	// CreateUserMod assigns a mod to a specific user.
	CreateUserMod(*model.UserModParams) error

	// DeleteUserMod removes a mod from a specific user.
	DeleteUserMod(*model.UserModParams) error

	// GetUserPacks retrieves packs for a user.
	GetUserPacks(*model.UserPackParams) (*model.UserPacks, error)

	// GetUserHasPack checks if a specific pack is assigned to a user.
	GetUserHasPack(*model.UserPackParams) bool

	// CreateUserPack assigns a pack to a specific user.
	CreateUserPack(*model.UserPackParams) error

	// DeleteUserPack removes a pack from a specific user.
	DeleteUserPack(*model.UserPackParams) error

	// GetUserTeams retrieves teams for a user.
	GetUserTeams(*model.UserTeamParams) (*model.TeamUsers, error)

	// GetUserHasTeam checks if a specific team is assigned to a user.
	GetUserHasTeam(*model.UserTeamParams) bool

	// CreateUserTeam assigns a team to a specific user.
	CreateUserTeam(*model.UserTeamParams) error

	// DeleteUserTeam removes a team from a specific user.
	DeleteUserTeam(*model.UserTeamParams) error

	// GetVersions retrieves all available versions from the database.
	GetVersions(int) (*model.Versions, error)

	// CreateVersion creates a new version.
	CreateVersion(int, *model.Version) error

	// UpdateVersion updates a version.
	UpdateVersion(int, *model.Version) error

	// DeleteVersion deletes a version.
	DeleteVersion(int, *model.Version) error

	// GetVersion retrieves a specific version from the database.
	GetVersion(int, string) (*model.Version, *gorm.DB)

	// GetVersionBuilds retrieves builds for a version.
	GetVersionBuilds(*model.VersionBuildParams) (*model.BuildVersions, error)

	// GetVersionHasBuild checks if a specific build is assigned to a version.
	GetVersionHasBuild(*model.VersionBuildParams) bool

	// CreateVersionBuild assigns a build to a specific version.
	CreateVersionBuild(*model.VersionBuildParams) error

	// DeleteVersionBuild removes a build from a specific version.
	DeleteVersionBuild(*model.VersionBuildParams) error

	// GetTeams retrieves all available teams from the database.
	GetTeams() (*model.Teams, error)

	// CreateTeam creates a new team.
	CreateTeam(*model.Team) error

	// UpdateTeam updates a team.
	UpdateTeam(*model.Team) error

	// DeleteTeam deletes a team.
	DeleteTeam(*model.Team) error

	// GetTeam retrieves a specific team from the database.
	GetTeam(string) (*model.Team, *gorm.DB)

	// GetTeamUsers retrieves users for a team.
	GetTeamUsers(*model.TeamUserParams) (*model.TeamUsers, error)

	// GetTeamHasUser checks if a specific user is assigned to a team.
	GetTeamHasUser(*model.TeamUserParams) bool

	// CreateTeamUser assigns a user to a specific team.
	CreateTeamUser(*model.TeamUserParams) error

	// DeleteTeamUser removes a user from a specific team.
	DeleteTeamUser(*model.TeamUserParams) error

	// GetTeamPacks retrieves packs for a team.
	GetTeamPacks(*model.TeamPackParams) (*model.TeamPacks, error)

	// GetTeamHasPack checks if a specific pack is assigned to a team.
	GetTeamHasPack(*model.TeamPackParams) bool

	// CreateTeamPack assigns a pack to a specific team.
	CreateTeamPack(*model.TeamPackParams) error

	// DeleteTeamPack removes a pack from a specific team.
	DeleteTeamPack(*model.TeamPackParams) error

	// GetTeamMods retrieves mods for a team.
	GetTeamMods(*model.TeamModParams) (*model.TeamMods, error)

	// GetTeamHasMod checks if a specific mod is assigned to a team.
	GetTeamHasMod(*model.TeamModParams) bool

	// CreateTeamMod assigns a mod to a specific team.
	CreateTeamMod(*model.TeamModParams) error

	// DeleteTeamMod removes a mod from a specific team.
	DeleteTeamMod(*model.TeamModParams) error

	// GetMinecrafts retrieves all available minecrafts from the database.
	GetMinecrafts() (*model.Minecrafts, error)

	// SyncMinecraft creates or updates a minecraft record.
	SyncMinecraft(*minecraft.Version) (*model.Minecraft, error)

	// GetMinecraft retrieves a specific minecraft from the database.
	GetMinecraft(string) (*model.Minecraft, *gorm.DB)

	// GetMinecraftBuilds retrieves builds for a minecraft.
	GetMinecraftBuilds(*model.MinecraftBuildParams) (*model.Builds, error)

	// GetMinecraftHasBuild checks if a specific build is assigned to a minecraft.
	GetMinecraftHasBuild(*model.MinecraftBuildParams) bool

	// CreateMinecraftBuild assigns a build to a specific minecraft.
	CreateMinecraftBuild(*model.MinecraftBuildParams) error

	// DeleteMinecraftBuild removes a build from a specific minecraft.
	DeleteMinecraftBuild(*model.MinecraftBuildParams) error

	// GetForges retrieves all available forges from the database.
	GetForges() (*model.Forges, error)

	// SyncForge creates or updates a forge record.
	SyncForge(*forge.Number) (*model.Forge, error)

	// GetForge retrieves a specific forge from the database.
	GetForge(string) (*model.Forge, *gorm.DB)

	// GetForgeBuilds retrieves builds for a forge.
	GetForgeBuilds(*model.ForgeBuildParams) (*model.Builds, error)

	// GetForgeHasBuild checks if a specific build is assigned to a forge.
	GetForgeHasBuild(*model.ForgeBuildParams) bool

	// CreateForgeBuild assigns a build to a specific forge.
	CreateForgeBuild(*model.ForgeBuildParams) error

	// DeleteForgeBuild removes a build from a specific forge.
	DeleteForgeBuild(*model.ForgeBuildParams) error

	GetSolderPacks() (*model.Packs, error)

	GetSolderPack(string, string) (*model.Pack, error)

	GetSolderBuild(string, string, string) (*model.Build, error)
}
