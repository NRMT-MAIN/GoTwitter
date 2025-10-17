-- +goose Up
-- +goose StatementBegin
CREATE TABLE tweet_tags (
	tweet_id BIGINT UNSIGNED NOT NULL,
	tag_id BIGINT UNSIGNED NOT NULL,
	PRIMARY KEY (tweet_id, tag_id),
	FOREIGN KEY (tweet_id) REFERENCES tweets(id),
	FOREIGN KEY (tag_id) REFERENCES tags(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tweet_tags if exists;
-- +goose StatementEnd
