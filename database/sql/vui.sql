-- name: ReadPreOrderChoice :many
SELECT message, seq_id
FROM pre_order_choice;

-- name: ReadProOrderChoice :many
SELECT seq_id, message, target
FROM pro_order_choice p, order_choice_target t
WHERE p.seq_id = t.target_id;

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


-- name: GetEntityInfo :many
SELECT e.typename, d.dinner_id, t.target_id, c.message
FROM dinner d, pro_order_choice c, order_choice_target t, entity_type e, 
WHERE d.entity_id = board_entity.entity_id
AND 

UNION ALL
