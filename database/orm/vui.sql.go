// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: vui.sql

package orm

import (
	"context"
)

const getCountId = `-- name: GetCountId :one
SELECT count_id
FROM pro_order_choice, entity_count
WHERE pro_order_choice.seq_id = entity_count.target_id
AND pro_order_choice.seq_id = (?)
`

func (q *Queries) GetCountId(ctx context.Context, seqID int32) (int32, error) {
	row := q.db.QueryRowContext(ctx, getCountId, seqID)
	var count_id int32
	err := row.Scan(&count_id)
	return count_id, err
}

const getDinnerId = `-- name: GetDinnerId :one
SELECT dinner_id
FROM pro_order_choice, board_entity, dinner
WHERE pro_order_choice.seq_id = (?)
AND pro_order_choice.seq_id = board_entity.target_id
AND board_entity.entity_id = dinner.entity_id
`

func (q *Queries) GetDinnerId(ctx context.Context, seqID int32) (int32, error) {
	row := q.db.QueryRowContext(ctx, getDinnerId, seqID)
	var dinner_id int32
	err := row.Scan(&dinner_id)
	return dinner_id, err
}

const getMenuId = `-- name: GetMenuId :one
SELECT menu_id
FROM pro_order_choice, board_entity, menu
WHERE pro_order_choice.seq_id = (?)
AND pro_order_choice.seq_id = board_entity.target_id
AND board_entity.entity_id = menu.entity_id
`

func (q *Queries) GetMenuId(ctx context.Context, seqID int32) (int32, error) {
	row := q.db.QueryRowContext(ctx, getMenuId, seqID)
	var menu_id int32
	err := row.Scan(&menu_id)
	return menu_id, err
}

const getOptionId = `-- name: GetOptionId :one
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
AND board_entity.entity_id = menu_option2.entity_id
`

type GetOptionIdParams struct {
	SeqID   int32
	SeqID_2 int32
}

func (q *Queries) GetOptionId(ctx context.Context, arg GetOptionIdParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, getOptionId, arg.SeqID, arg.SeqID_2)
	var option_id int32
	err := row.Scan(&option_id)
	return option_id, err
}

const getStyleId = `-- name: GetStyleId :one
SELECT style_id
FROM pro_order_choice, board_entity, style
WHERE pro_order_choice.seq_id = (?)
AND pro_order_choice.seq_id = board_entity.target_id
AND board_entity.entity_id = style.entity_id
`

func (q *Queries) GetStyleId(ctx context.Context, seqID int32) (int32, error) {
	row := q.db.QueryRowContext(ctx, getStyleId, seqID)
	var style_id int32
	err := row.Scan(&style_id)
	return style_id, err
}

const readPreOrderChoice = `-- name: ReadPreOrderChoice :many
SELECT seq_id, message
FROM pre_order_choice
`

type ReadPreOrderChoiceRow struct {
	SeqID   int32
	Message string
}

func (q *Queries) ReadPreOrderChoice(ctx context.Context) ([]ReadPreOrderChoiceRow, error) {
	rows, err := q.db.QueryContext(ctx, readPreOrderChoice)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ReadPreOrderChoiceRow
	for rows.Next() {
		var i ReadPreOrderChoiceRow
		if err := rows.Scan(&i.SeqID, &i.Message); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const readPreOrderChoiceNxtSeq = `-- name: ReadPreOrderChoiceNxtSeq :many
SELECT seq_id, nxt_id
FROM pre_order_choice_nxt_seq
`

func (q *Queries) ReadPreOrderChoiceNxtSeq(ctx context.Context) ([]PreOrderChoiceNxtSeq, error) {
	rows, err := q.db.QueryContext(ctx, readPreOrderChoiceNxtSeq)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PreOrderChoiceNxtSeq
	for rows.Next() {
		var i PreOrderChoiceNxtSeq
		if err := rows.Scan(&i.SeqID, &i.NxtID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const readProOrderChoice = `-- name: ReadProOrderChoice :many
SELECT seq_id, message, target, typename
FROM pro_order_choice, entity_type
WHERE entity_type.type_id = pro_order_choice.type_id
`

type ReadProOrderChoiceRow struct {
	SeqID    int32
	Message  string
	Target   string
	Typename string
}

func (q *Queries) ReadProOrderChoice(ctx context.Context) ([]ReadProOrderChoiceRow, error) {
	rows, err := q.db.QueryContext(ctx, readProOrderChoice)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ReadProOrderChoiceRow
	for rows.Next() {
		var i ReadProOrderChoiceRow
		if err := rows.Scan(
			&i.SeqID,
			&i.Message,
			&i.Target,
			&i.Typename,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const readProOrderChoiceNxtSeq = `-- name: ReadProOrderChoiceNxtSeq :many
SELECT seq_id, nxt_id
FROM pro_order_choice_nxt_seq
`

func (q *Queries) ReadProOrderChoiceNxtSeq(ctx context.Context) ([]ProOrderChoiceNxtSeq, error) {
	rows, err := q.db.QueryContext(ctx, readProOrderChoiceNxtSeq)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProOrderChoiceNxtSeq
	for rows.Next() {
		var i ProOrderChoiceNxtSeq
		if err := rows.Scan(&i.SeqID, &i.NxtID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
