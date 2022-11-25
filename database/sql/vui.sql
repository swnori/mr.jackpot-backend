-- name: ReadPreOrderChoice :many
SELECT message, seq_id
FROM pre_order_choice;

-- name: ReadProOrderChoice :many
SELECT seq_id, message, target
FROM pro_order_choice;

-- name: ReadPreOrderChoiceNxtSeq :many
SELECT seq_id, nxt_id
FROM pre_order_choice_nxt_seq;

-- name: ReadProOrderChoiceNxtSeq :many
SELECT seq_id, nxt_id
FROM pro_order_choice_nxt_seq;

-- name: GetDinnerEntity :many
SELECT target_id, dinner_id, typename
FROM  dinner, entity_type, board_entity
WHERE dinner.entity_id = board_entity.entity_id
AND entity_type.type_id = board_entity.type_id;

-- name: GetMenuEntity :many
SELECT target_id, menu_id, typename
FROM  menu, entity_type, board_entity
WHERE menu.entity_id = board_entity.entity_id
AND entity_type.type_id = board_entity.type_id;

-- name: GetStyleEntity :many
SELECT target_id, style_id, typename
FROM  style, entity_type, board_entity
WHERE style.entity_id = board_entity.entity_id
AND entity_type.type_id = board_entity.type_id;

-- name: GetOptionEntity :many
SELECT target_id, option_id, typename
FROM  menu_option1, entity_type, board_entity
WHERE menu_option1.entity_id = board_entity.entity_id
AND entity_type.type_id = board_entity.type_id
UNION
SELECT target_id, option_id, typename
FROM  menu_option2, entity_type, board_entity
WHERE menu_option2.entity_id = board_entity.entity_id
AND entity_type.type_id = board_entity.type_id;

-- name: GetAllEntityIdList :many
SELECT seq_id as target_id
FROM pro_order_choice;

-- name: Ping :exec
INSERT INTO user
VALUES ();


