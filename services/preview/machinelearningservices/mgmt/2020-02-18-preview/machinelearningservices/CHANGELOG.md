Generated from https://github.com/Azure/azure-rest-api-specs/tree/3c764635e7d442b3e74caf593029fcd440b3ef82

Code generator @microsoft.azure/autorest.go@~2.1.161

## Breaking Changes

- Function `NewListUsagesResultPage` parameter(s) have been changed from `(func(context.Context, ListUsagesResult) (ListUsagesResult, error))` to `(ListUsagesResult, func(context.Context, ListUsagesResult) (ListUsagesResult, error))`
- Function `NewListAmlUserFeatureResultPage` parameter(s) have been changed from `(func(context.Context, ListAmlUserFeatureResult) (ListAmlUserFeatureResult, error))` to `(ListAmlUserFeatureResult, func(context.Context, ListAmlUserFeatureResult) (ListAmlUserFeatureResult, error))`
- Function `NewListWorkspaceQuotasPage` parameter(s) have been changed from `(func(context.Context, ListWorkspaceQuotas) (ListWorkspaceQuotas, error))` to `(ListWorkspaceQuotas, func(context.Context, ListWorkspaceQuotas) (ListWorkspaceQuotas, error))`
- Function `NewWorkspaceListResultPage` parameter(s) have been changed from `(func(context.Context, WorkspaceListResult) (WorkspaceListResult, error))` to `(WorkspaceListResult, func(context.Context, WorkspaceListResult) (WorkspaceListResult, error))`
- Function `NewSkuListResultPage` parameter(s) have been changed from `(func(context.Context, SkuListResult) (SkuListResult, error))` to `(SkuListResult, func(context.Context, SkuListResult) (SkuListResult, error))`
- Function `NewPaginatedComputeResourcesListPage` parameter(s) have been changed from `(func(context.Context, PaginatedComputeResourcesList) (PaginatedComputeResourcesList, error))` to `(PaginatedComputeResourcesList, func(context.Context, PaginatedComputeResourcesList) (PaginatedComputeResourcesList, error))`

## New Content

- New const `ComputeInstanceAuthorizationTypePersonal`
- New function `PossibleComputeInstanceAuthorizationTypeValues() []ComputeInstanceAuthorizationType`
- New struct `AssignedUser`
- New struct `PersonalComputeInstanceSettings`
- New field `PersonalComputeInstanceSettings` in struct `ComputeInstanceProperties`
- New field `ComputeInstanceAuthorizationType` in struct `ComputeInstanceProperties`
