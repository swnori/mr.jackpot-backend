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

-- name: GetDinnerId :one
SELECT dinner_id
FROM dinner
WHERE entity_id = (?);

-- name: GetMenuId :one
SELECT menu_id
FROM menu
WHERE entity_id = (?);

-- name: GetOption1Id :one
SELECT option_id
FROM menu_option1
WHERE entity_id = (?);

-- name: GetOption2Id :one
SELECT option_id
FROM menu_option2
WHERE entity_id = (?);

