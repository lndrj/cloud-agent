package repository

import (
	"collector/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MetricsRepository struct {
	pool *pgxpool.Pool
}

func NewMetricsRepository(pool *pgxpool.Pool) *MetricsRepository {
	return &MetricsRepository{pool: pool}
}

func (r *MetricsRepository) Save(
	ctx context.Context,
	agentID string,
	metrics models.AgentMetrics,
) error {

	insertQuery := `
		INSERT INTO agent_metrics (
			agent_id,
			hostname,
			uptime,
			memory_available_mb,
			memory_used_mb,
			cpu_percent,
			disk_used_percent
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.pool.Exec(
		ctx,
		insertQuery,
		agentID,
		metrics.Hostname,
		metrics.Uptime,
		metrics.MemoryAvailableMb,
		metrics.MemoryUsedMb,
		metrics.CpuPercent,
		metrics.DiskUsedPercent,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *MetricsRepository) GetHistory(
	ctx context.Context,
	agentID string,
	limit int,
) ([]models.MetricsSnapshot, error) {

	selectQuery := `
		SELECT
			agent_id,
			hostname,
			uptime,
			memory_available_mb,
			memory_used_mb,
			cpu_percent,
			disk_used_percent,
			collected_at
		FROM agent_metrics
		WHERE agent_id = $1
		ORDER BY collected_at DESC
		LIMIT $2
	`

	rows, err := r.pool.Query(
		ctx,
		selectQuery,
		agentID,
		limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []models.MetricsSnapshot

	for rows.Next() {
		var snapshot models.MetricsSnapshot

		err := rows.Scan(
			&snapshot.AgentId,
			&snapshot.Metrics.Hostname,
			&snapshot.Metrics.Uptime,
			&snapshot.Metrics.MemoryAvailableMb,
			&snapshot.Metrics.MemoryUsedMb,
			&snapshot.Metrics.CpuPercent,
			&snapshot.Metrics.DiskUsedPercent,
			&snapshot.CollectedAt,
		)
		if err != nil {
			return nil, err
		}

		history = append(history, snapshot)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return history, nil
}
