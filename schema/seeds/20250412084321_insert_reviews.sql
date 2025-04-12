-- +goose Up
-- +goose StatementBegin
INSERT INTO reviews (reviewer_id, target_id,booking_id, rating, comment) VALUES ('c72ff794-3d1d-4583-a988-eb89fe509ee3', 'c72ff794-3d1d-4583-a988-eb89fe509ee3', 'c72ff794-3d1d-4583-a988-eb89fe509ee3', 4, 'new comment');
INSERT INTO reviews (reviewer_id, target_id,booking_id, rating, comment) VALUES ('c72ff794-3d1d-4583-a988-eb89fe509ee3', 'c72ff794-3d1d-4583-a988-eb89fe509ee3', 'c72ff794-3d1d-4583-a988-eb89fe509ee3', 4, 'new comment');
INSERT INTO reviews (reviewer_id, target_id,booking_id, rating, comment) VALUES ('c72ff794-3d1d-4583-a988-eb89fe509ee3', 'c72ff794-3d1d-4583-a988-eb89fe509ee3', 'c72ff794-3d1d-4583-a988-eb89fe509ee3', 4, 'new comment');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE reviews CASCADE;
-- +goose StatementEnd
