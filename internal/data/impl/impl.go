package impl

import "github.com/google/wire"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewUserRepository,
	NewArticleRepository,
	NewCommentRepository,
	NewFavoriteRepository,
	NewTagRepository,
)
