-- name: ReadDinnerEntity :many
SELECT dinner_id, name, price
FROM dinner d, board_entity e
WHERE d.entity_id = e.entity_id;

-- name: ReadDinnersMenu :many
SELECT menu_id, default_count
FROM dinners_menu
WHERE dinner_id = (?);

-- name: ReadMenuEntity :many
SELECT menu_id, e.name, price, option1_name, option2_name, t.name as typename
FROM menu m, board_entity e, menu_type t
WHERE m.entity_id = e.entity_id
  AND t.id = m.type_id;

-- name: ReadOption1Entity :many
SELECT option_id, name, price
FROM menu_option1 o, board_entity e
WHERE o.entity_id = e.entity_id
  AND o.menu_id = (?);

-- name: ReadOption2Entity :many
SELECT option_id, name, price
FROM menu_option2 o, board_entity e
WHERE o.entity_id = e.entity_id
  AND o.menu_id = (?);

-- name: ReadStyleEntity :many
SELECT style_id, name, description, price
FROM style s, board_entity e
WHERE s.entity_id = e.entity_id;

-- name: ReadOrderState :many
SELECT state_id, name
FROM state;