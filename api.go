package coreapiinterfaces

import "github.com/gin-gonic/gin"

func NodeApiBinding(router *gin.Engine, api NodeApi) {
	router.GET("/nodes/:id", api.GetNode)
	router.DELETE("/nodes/:id", api.DeleteNode)
	router.POST("/nodes", api.CreateNode)
	router.GET("/nodes", api.ListNode)
	router.GET("/nodes/:id/config", api.GetNodeConfig)
	router.PATCH("/nodes/:id/config", api.UpdateNodeConfig)
	router.GET("/nodes/config/api", api.GetNodeApiConfigByDomain)
}
func RoleApiBinding(router *gin.Engine, api RoleApi) {
	router.PUT("/roles/:id", api.UpdateRole)
	router.DELETE("/roles/:id", api.DeleteRole)
	router.GET("/roles/:id", api.GetRole)
	router.GET("/roles", api.ListRoles)
	router.POST("/roles", api.CreateRole)
	router.GET("/roles/:id/permissions", api.GetRolePermissions)
	router.PUT("/roles/:id/permissions", api.UpdateRolePermissions)
}
func AdminApiBinding(router *gin.Engine, api AdminApi) {
	router.PUT("/admins/:id/password", api.UpdateAdminPassword)
	router.PUT("/admins/:id/roles", api.UpdateAdminRole)
	router.POST("/admins", api.CreateAdmin)
	router.GET("/admins", api.ListAdmin)
	router.PUT("/admins/self/password", api.UpdateAdminSelfPassword)
	router.GET("/admins/self/totp", api.GetTotp)
	router.PUT("/admins/self/totp", api.UpdateTotp)
	router.GET("/admins/:id", api.GetAdmin)
	router.DELETE("/admins/:id", api.DeleteAdmin)
	router.GET("/admins/self", api.GetAdminSelf)
}
func SessionApiBinding(router *gin.Engine, api SessionApi) {
	router.POST("/sessions", api.CreateSession)
}

type AdminApi interface {
	UpdateAdminPassword(ctx *gin.Context)
	UpdateAdminRole(ctx *gin.Context)
	CreateAdmin(ctx *gin.Context)
	ListAdmin(ctx *gin.Context)
	UpdateAdminSelfPassword(ctx *gin.Context)
	GetTotp(ctx *gin.Context)
	UpdateTotp(ctx *gin.Context)
	GetAdmin(ctx *gin.Context)
	DeleteAdmin(ctx *gin.Context)
	GetAdminSelf(ctx *gin.Context)
}
type SessionApi interface {
	CreateSession(ctx *gin.Context)
}
type NodeApi interface {
	GetNode(ctx *gin.Context)
	DeleteNode(ctx *gin.Context)
	CreateNode(ctx *gin.Context)
	ListNode(ctx *gin.Context)
	GetNodeConfig(ctx *gin.Context)
	UpdateNodeConfig(ctx *gin.Context)
	GetNodeApiConfigByDomain(ctx *gin.Context)
}
type RoleApi interface {
	UpdateRole(ctx *gin.Context)
	DeleteRole(ctx *gin.Context)
	GetRole(ctx *gin.Context)
	ListRoles(ctx *gin.Context)
	CreateRole(ctx *gin.Context)
	GetRolePermissions(ctx *gin.Context)
	UpdateRolePermissions(ctx *gin.Context)
}
type CreateAdminRequest struct {
	Account  string  `json:"account" xml:"account"`
	Password string  `json:"password" xml:"password"`
	Role     *string `json:"role" xml:"role"`
}
type Role struct {
	Description string `json:"description" xml:"description"`
	Id          string `json:"id" xml:"id"`
	Name        string `json:"name" xml:"name"`
}
type Ordering string

const ASCENDING Ordering = "ASCENDING"
const DESCENDING Ordering = "DESCENDING"

type Node struct {
	UpdatedAt   int64   `json:"updatedAt" xml:"updatedAt"`
	Id          string  `json:"id" xml:"id"`
	Name        string  `json:"name" xml:"name"`
	SecurityKey *string `json:"securityKey" xml:"securityKey"`
	Description *string `json:"description" xml:"description"`
	CreatedAt   int64   `json:"createdAt" xml:"createdAt"`
}
type Admin struct {
	Id          string       `json:"id" xml:"id"`
	Account     string       `json:"account" xml:"account"`
	PhoneNumber *PhoneNumber `json:"phoneNumber" xml:"phoneNumber"`
	Role        *string      `json:"role" xml:"role"`
	HasTotp     bool         `json:"hasTotp" xml:"hasTotp"`
	CreatedAt   int64        `json:"createdAt" xml:"createdAt"`
	UpdatedAt   int64        `json:"updatedAt" xml:"updatedAt"`
}
type Totp struct {
	Url string `json:"url" xml:"url"`
}
type UpdateRoleRequest struct {
	Name        string `json:"name" xml:"name"`
	Description string `json:"description" xml:"description"`
}
type CreateNodeRequest struct {
	Name        *string `json:"name" xml:"name"`
	Description *string `json:"description" xml:"description"`
}
type ApiConfig struct {
	MerchantUrl *string `json:"merchantUrl" xml:"merchantUrl"`
}
type PhoneNumber struct {
	CountryCode string `json:"countryCode" xml:"countryCode"`
	Number      string `json:"number" xml:"number"`
}
type PermissionDomain struct {
	Domain      string    `json:"domain" xml:"domain"`
	Description string    `json:"description" xml:"description"`
	Permissions *[]string `json:"permissions" xml:"permissions"`
}
type RedisConfig struct {
	Port     *string `json:"port" xml:"port"`
	Password *string `json:"password" xml:"password"`
	Host     *string `json:"host" xml:"host"`
}
type DBConfig struct {
	Username string        `json:"username" xml:"username"`
	Password string        `json:"password" xml:"password"`
	Database string        `json:"database" xml:"database"`
	Type     *DatabaseType `json:"type" xml:"type"`
	Host     string        `json:"host" xml:"host"`
	Port     string        `json:"port" xml:"port"`
}
type CreateRoleRequest struct {
	Name        string `json:"name" xml:"name"`
	Description string `json:"description" xml:"description"`
}
type RoleList struct {
	Items      []Role     `json:"items" xml:"items"`
	Pagination Pagination `json:"pagination" xml:"pagination"`
}
type Pagination struct {
	Index int64 `json:"index" xml:"index"`
	Limit int64 `json:"limit" xml:"limit"`
	Total int64 `json:"total" xml:"total"`
}
type NodeConfig struct {
	TencentCloud   *TencentCloudConfig `json:"tencentCloud" xml:"tencentCloud"`
	FrontendDomain *string             `json:"frontendDomain" xml:"frontendDomain"`
	Api            *ApiConfig          `json:"api" xml:"api"`
	Server         *ServerConfig       `json:"server" xml:"server"`
	Database       *DBConfig           `json:"database" xml:"database"`
	Redis          *RedisConfig        `json:"redis" xml:"redis"`
}
type AdminList struct {
	Items      []Admin    `json:"items" xml:"items"`
	Pagination Pagination `json:"pagination" xml:"pagination"`
}
type Session struct {
	Token string `json:"token" xml:"token"`
}
type AddRolesRequest struct {
	RolesIds *[]string `json:"rolesIds" xml:"rolesIds"`
}
type UpdatePasswordRequest struct {
	Password    string  `json:"password" xml:"password"`
	OldPassword *string `json:"oldPassword" xml:"oldPassword"`
}
type CreateSessionRequest struct {
	Account  *string `json:"account" xml:"account"`
	Password *string `json:"password" xml:"password"`
}
type UpdateTotpRequest struct {
	Url      string `json:"url" xml:"url"`
	TotpCode string `json:"totpCode" xml:"totpCode"`
}
type TencentCloudConfig struct {
	SecretId  *string `json:"secretId" xml:"secretId"`
	SecretKey *string `json:"secretKey" xml:"secretKey"`
}
type DatabaseType string

const MYSQL DatabaseType = "MYSQL"
const POSTGRES DatabaseType = "POSTGRES"

type NodeList struct {
	Items      []Node     `json:"items" xml:"items"`
	Pagination Pagination `json:"pagination" xml:"pagination"`
}
type ServerConfig struct {
	Port      *string `json:"port" xml:"port"`
	JwtSecret *string `json:"jwtSecret" xml:"jwtSecret"`
}
