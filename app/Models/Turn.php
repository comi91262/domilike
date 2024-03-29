<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Turn extends Model
{
    protected $fillable = [
        'id', 'user_id', 'is_turn',
    ];

    private $MEMBER_COUNT = 2;

    /**
     * ターンテーブル（順番を決めるテーブル)を初期化する
     */
    public function init($memberCount)
    {
        $this->truncate();

        $members = [];

        for ($i = 1; $i <= $memberCount; $i++) {
            $members[] = $i;
        }

        shuffle($members);

        foreach ($members as $member) {
            $turn = new Turn(); //foreachのたびにnewしないとsaveされないみたい
            $turn->user_id = $member;
            $turn->save();
        }

        //初めのレコードのプレイヤーをスタートプレイヤーとする
        //既存のレコードを更新する場合、以下のようにする
        $turn = Turn::find(1);
        $turn->is_turn = true;
        $turn->save();
    }

    /**
     * ターンテーブル（順番を決めるテーブル)から自身のプレイヤーデータを引き出す
     */
    public function getPlayer($id)
    {
        $player = $this->where('user_id', $id)->first();

        return $player;
    }

    /**
     *  次のターンのプレイヤーを決めるメソッド
     *  現在のプレイヤーのis_turnをfalseにし、
     *  次のプレイヤーのis_turnをtrueにする。
     */
    public function change($player_id)
    {
        $player = $this->where('user_id', $player_id)->first();
        $player->is_turn = false;
        $player->save();
        $next_id = $player->id % $this->MEMBER_COUNT + 1;

        $turn = Turn::find($next_id);
        $turn->is_turn = true;
        $turn->save();
    }

    /**
     *   参加者全てのidを取得する
     */
    public function getEntries()
    {
        $ids = [];
        $members = $this->all();

        foreach ($members as $member) {
            $ids[] = $member->user_id;
        }

        return $ids;
    }

    /**
     *   IDをシャッフルし、登録する
     */
    public function register($ids)
    {
        $this->truncate();
        shuffle($ids);

        foreach ($ids as $id) {
            $turn = new Turn(); //foreachのたびにnewしないとsaveされないみたい
            $turn->user_id = $id;
            $turn->save();
        }
    }

    /**
     *  　次のターンのユーザーIDを求めるメソッド
     */
    public function next($id)
    {
        $entryIds = $this->getEntries();

        $memberN = count($entryIds);
        $next = 0;

        for ($i = 0; $i < $memberN; $i++) {
            if ($id == $entryIds[$i]) {
                if ($i == $memberN - 1) {
                    $next = 0;
                } else {
                    $next = $i + 1;
                }
                break;
            }
        }

        return $this->find($next + 1)->user_id;
    }
}
