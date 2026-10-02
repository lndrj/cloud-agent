-- +goose Up

CREATE TABLE agent_metrics
(
    id                  BIGSERIAL PRIMARY KEY,
    agent_id            TEXT             NOT NULL,
    hostname            TEXT             NOT NULL,
    uptime              BIGINT           NOT NULL,
    memory_available_mb BIGINT           NOT NULL,
    memory_used_mb      BIGINT           NOT NULL,
    cpu_percent         DOUBLE PRECISION NOT NULL,
    disk_used_percent   DOUBLE PRECISION NOT NULL,
    collected_at        TIMESTAMPTZ      NOT NULL DEFAULT NOW()
);

-- +goose Down

DROP TABLE agent_metrics;