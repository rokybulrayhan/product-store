package lib

import (
	"context"
	"strconv"

	"github.com/go-contact-service/entity"
	"github.com/gosimple/slug"
)

type slugRepository interface {
	SearchSlug(context.Context, string) ([]string, error)
	SlugChecker(context.Context, string) (bool, error)
}

func GenerateSlugGlobal(ctx context.Context, slugText, name string, repo slugRepository) (string, error) {
	if slugText != "" {
		existingSlugs, err := repo.SearchSlug(ctx, slugText)
		if err != nil {
			return "", entity.ErrSlugGenerationFailed
		}
		if len(existingSlugs) > 0 {
			return "", entity.ErrSlugExists
		}
	}
	if slugText == "" {
		slugText = slug.MakeLang(name, "en")
	}
	existingSlugs, err := repo.SearchSlug(ctx, slugText)
	if err != nil {
		return "", entity.ErrSlugGenerationFailed
	}
	if len(existingSlugs) > 0 {
		slugText = slugText + "-" + strconv.Itoa(len(existingSlugs))
	}
	for _, sg := range existingSlugs {
		if sg == slugText {
			return "", entity.ErrSlugExists
		}
	}
	return slugText, nil
}
