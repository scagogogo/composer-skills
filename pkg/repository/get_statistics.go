package repository

import (
	"context"
	"fmt"

	"github.com/scagogogo/composer-crawler/pkg/domain"
)

// Statistics 获取仓库中的组件下载统计信息
func (x *Repository) Statistics(ctx context.Context) (*domain.StatisticsResponse, error) {
	targetUrl := fmt.Sprintf("%s/statistics.json", x.options.ServerUrl)
	return getJson[*domain.StatisticsResponse](ctx, x, targetUrl)
}
