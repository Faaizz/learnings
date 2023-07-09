from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class void(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class TodoItem(_message.Message):
    __slots__ = ["id", "text"]
    ID_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    id: int
    text: str
    def __init__(self, id: _Optional[int] = ..., text: _Optional[str] = ...) -> None: ...

class TodoItems(_message.Message):
    __slots__ = ["items"]
    ITEMS_FIELD_NUMBER: _ClassVar[int]
    items: _containers.RepeatedCompositeFieldContainer[TodoItem]
    def __init__(self, items: _Optional[_Iterable[_Union[TodoItem, _Mapping]]] = ...) -> None: ...
