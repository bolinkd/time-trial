
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO "Group2" (parent_id, name, organization_id, created_at, updated_at, tmp_club_id)
    ( SELECT null, c.name, c.organization_id, c.created_at, c.updated_at, c.id FROM club c );

INSERT INTO "Group2" (parent_id, name, organization_id, created_at, updated_at)
    ( SELECT
          (SELECT g2.id FROM "Group2" g2 WHERE g2.tmp_club_id = g.club_id),
          g.name,
          (SELECT g2.organization_id FROM "Group2" g2 WHERE g2.tmp_club_id = g.club_id),
          g.created_at,
          g.updated_at
      FROM "Group" as g);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

INSERT INTO club (id, name, abbreviation, organization_id, created_at, updated_at)
(
    SELECT g2.tmp_club_id, g2.name, null, g2.organization_id, g2.created_at, g2.updated_at
    FROM "Group2" g2
    WHERE g2.tmp_club_id IS NOT NULL
);

INSERT INTO "Group" (name, club_id, created_at, updated_at)
    (
        SELECT g2.name, (
            SELECT c.id FROM club c WHERE c.id =
                (SELECT gg.tmp_club_id FROM "Group2" gg WHERE gg.id = g2.parent_id)
        ), g2.created_at, g2.updated_at
        FROM "Group2" g2
        WHERE g2.tmp_club_id IS NULL
    );