-- name: ReadDinnerEntity :many
SELECT dinner_id, name, price
FROM dinner d, board_entity e
WHERE d.entity_id = e.entity_id;

-- name: ReadDinnersMenu :many
SELECT menu_id
FROM dinners_menu
WHERE dinner_id = (?);

-- name: ReadMenuEntity :many
SELECT menu_id, name, price, option1_name, option2_name
FROM menu m, board_entity e
WHERE n.entity_id = e.entity_id;

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
SELECT style_id, name, price
FROM style s, board_entity e
WHERE s.entity_id = e.entity_id;
