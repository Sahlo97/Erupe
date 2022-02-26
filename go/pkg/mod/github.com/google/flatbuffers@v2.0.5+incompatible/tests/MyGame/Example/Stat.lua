--[[ MyGame.Example.Stat

  Automatically generated by the FlatBuffers compiler, do not modify.
  Or modify. I'm a message, not a cop.

  flatc version: 2.0.5

  Declared by  : //monster_test.fbs
  Rooting type : MyGame.Example.Monster (//monster_test.fbs)

--]]

local flatbuffers = require('flatbuffers')

local Stat = {}
local mt = {}

function Stat.New()
  local o = {}
  setmetatable(o, {__index = mt})
  return o
end

function mt:Init(buf, pos)
  self.view = flatbuffers.view.New(buf, pos)
end

function mt:Id()
  local o = self.view:Offset(4)
  if o ~= 0 then
    return self.view:String(self.view.pos + o)
  end
end

function mt:Val()
  local o = self.view:Offset(6)
  if o ~= 0 then
    return self.view:Get(flatbuffers.N.Int64, self.view.pos + o)
  end
  return 0
end

function mt:Count()
  local o = self.view:Offset(8)
  if o ~= 0 then
    return self.view:Get(flatbuffers.N.Uint16, self.view.pos + o)
  end
  return 0
end

function Stat.Start(builder)
  builder:StartObject(3)
end

function Stat.AddId(builder, id)
  builder:PrependUOffsetTRelativeSlot(0, id, 0)
end

function Stat.AddVal(builder, val)
  builder:PrependInt64Slot(1, val, 0)
end

function Stat.AddCount(builder, count)
  builder:PrependUint16Slot(2, count, 0)
end

function Stat.End(builder)
  return builder:EndObject()
end

return Stat