-- name: ReadPreOrderChoice :many
SELECT seq_id, message
FROM pre_order_choice;

-- name: ReadProOrderChoice :many
SELECT seq_id, message, target, typename
FROM pro_order_choice, entity_type
WHERE entity_type.type_id = pro_order_choice.type_id;

-- name: ReadPreOrderChoiceNxtSeq :many
SELECT seq_id, nxt_id
FROM pre_order_choice_nxt_seq;

-- name: ReadProOrderChoiceNxtSeq :many
SELECT seq_id, nxt_id
FROM pro_order_choice_nxt_seq;

-- name: GetDinnerId :one
SELECT dinner_id
FROM pro_order_choice, board_entity, dinner
WHERE pro_order_choice.seq_id = (?)
AND pro_order_choice.seq_id = board_entity.target_id
AND board_entity.entity_id = dinner.entity_id;

-- name: GetMenuId :one
SELECT menu_id
FROM pro_order_choice, board_entity, menu
WHERE pro_order_choice.seq_id = (?)
AND pro_order_choice.seq_id = board_entity.target_id
AND board_entity.entity_id = menu.entity_id;

-- name: GetOptionId :one
SELECT option_id
FROM pro_order_choice, board_entity, menu_option1
WHERE pro_order_choice.seq_id = (?)
AND pro_order_choice.seq_id = board_entity.target_id
AND board_entity.entity_id = menu_option1.entity_id
UNION ALL
SELECT option_id
FROM pro_order_choice, board_entity, menu_option2
WHERE pro_order_choice.seq_id = (?)
AND pro_order_choice.seq_id = board_entity.target_id
AND board_entity.entity_id = menu_option2.entity_id;

-- name: GetStyleId :one
SELECT style_id
FROM pro_order_choice, board_entity, style
WHERE pro_order_choice.seq_id = (?)
AND pro_order_choice.seq_id = board_entity.target_id
AND board_entity.entity_id = style.entity_id;

-- name: GetCountId :one
SELECT count_id
FROM pro_order_choice, entity_count
WHERE pro_order_choice.seq_id = entity_count.target_id
AND pro_order_choice.seq_id = (?);

