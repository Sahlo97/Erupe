--[[ MyGame.Example.Color

  Automatically generated by the FlatBuffers compiler, do not modify.
  Or modify. I'm a message, not a cop.

  flatc version: 2.0.5

  Declared by  : //monster_test.fbs
  Rooting type : MyGame.Example.Monster (//monster_test.fbs)

--]]

-- Composite components of Monster color.
local Color = {
  Red = 1,
  -- \brief color Green
  -- Green is bit_flag with value (1u << 1)
  Green = 2,
  -- \brief color Blue (1u << 3)
  Blue = 8,
}

return Color