<?php
// automatically generated by the FlatBuffers compiler, do not modify

namespace NamespaceA\NamespaceB;

class UnionInNestedNS
{
    const NONE = 0;
    const TableInNestedNS = 1;

    private static $names = array(
        UnionInNestedNS::NONE=>"NONE",
        UnionInNestedNS::TableInNestedNS=>"TableInNestedNS",
    );

    public static function Name($e)
    {
        if (!isset(self::$names[$e])) {
            throw new \Exception();
        }
        return self::$names[$e];
    }
}
