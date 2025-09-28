package services

// RoleService provides business logic for managing roles, permissions, and user-role assignments.
// It acts as an intermediary between the controllers and repositories, handling validation and orchestration.

import (
	repositories "AuthInGo/db/repositories"
	"AuthInGo/models"
)

type RoleService interface {
	GetRoleById(id int64) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRole(name string, description string) (*models.Role, error)
	DeleteRoleById(id int64) error
	UpdateRole(id int64, name string, description string) (*models.Role, error)
	GetRolePermissions(roleId int64) ([]*models.RolePermission, error)
	AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error)
	RemovePermissionFromRole(roleId int64, permissionId int64) error
	GetAllRolePermissions() ([]*models.RolePermission, error)
	AssignRoleToUser(userId int64, roleId int64) error
}

type RoleServiceImpl struct {
	roleRepository           repositories.RoleRepository
	rolePermissionRepository repositories.RolePermissionRepository
	userRoleRepository       repositories.UserRoleRepository
}

// NewRoleService creates a new instance of RoleServiceImpl with the provided repositories.
func NewRoleService(roleRepo repositories.RoleRepository, rolePermissionRepo repositories.RolePermissionRepository, userRoleRepo repositories.UserRoleRepository) RoleService {
	return &RoleServiceImpl{
		roleRepository:           roleRepo,
		rolePermissionRepository: rolePermissionRepo,
		userRoleRepository:       userRoleRepo,
	}
}

// GetRoleById retrieves a role by its ID.
func (s *RoleServiceImpl) GetRoleById(id int64) (*models.Role, error) {
	return s.roleRepository.GetRoleById(id)
}

// GetRoleByName retrieves a role by its name.
func (s *RoleServiceImpl) GetRoleByName(name string) (*models.Role, error) {
	return s.roleRepository.GetRoleByName(name)
}

// GetAllRoles retrieves all roles.
func (s *RoleServiceImpl) GetAllRoles() ([]*models.Role, error) {
	return s.roleRepository.GetAllRoles()
}

// CreateRole creates a new role with the given name and description.
func (s *RoleServiceImpl) CreateRole(name string, description string) (*models.Role, error) {
	return s.roleRepository.CreateRole(name, description)
}

// DeleteRoleById deletes a role by its ID.
func (s *RoleServiceImpl) DeleteRoleById(id int64) error {
	return s.roleRepository.DeleteRoleById(id)
}

// UpdateRole updates a role's name and description by ID.
func (s *RoleServiceImpl) UpdateRole(id int64, name string, description string) (*models.Role, error) {
	return s.roleRepository.UpdateRole(id, name, description)
}

// GetRolePermissions retrieves all permissions for a given role.
func (s *RoleServiceImpl) GetRolePermissions(roleId int64) ([]*models.RolePermission, error) {
	return s.rolePermissionRepository.GetRolePermissionByRoleId(roleId)
}

// AddPermissionToRole adds a permission to a role.
func (s *RoleServiceImpl) AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error) {
	return s.rolePermissionRepository.AddPermissionToRole(roleId, permissionId)
}

// RemovePermissionFromRole removes a permission from a role.
func (s *RoleServiceImpl) RemovePermissionFromRole(roleId int64, permissionId int64) error {
	return s.rolePermissionRepository.RemovePermissionFromRole(roleId, permissionId)
}

// GetAllRolePermissions retrieves all role-permission associations.
func (s *RoleServiceImpl) GetAllRolePermissions() ([]*models.RolePermission, error) {
	return s.rolePermissionRepository.GetAllRolePermissions()
}

// AssignRoleToUser assigns a role to a user.
func (s *RoleServiceImpl) AssignRoleToUser(userId int64, roleId int64) error {
	return s.userRoleRepository.AssignRoleToUser(userId, roleId)
}
