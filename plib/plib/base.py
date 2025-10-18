import json

class Point:

    def __init__(self, x: int, y: int) -> None:
        if not isinstance(x, int) or not isinstance(y, int):
            raise TypeError('x, y should be of int type')
        self.x = x
        self.y = y

    def __eq__(self, other_point: 'Point') -> bool:
        return self.x == other_point.x and self.y == other_point.y
    
    def __add__(self, other_point: 'Point') -> 'Point':
        return Point(self.x + other_point.x, self.y + other_point.y)
    
    def __iadd__(self, other_point: 'Point') -> 'Point':
        return Point(self.x + other_point.x, self.y + other_point.y)
    
    def __sub__(self, other_point: 'Point') -> 'Point':
        return Point(self.x - other_point.x, self.y - other_point.y)
    
    def __isub__(self, other_point: 'Point') -> 'Point':
        return Point(self.x - other_point.x, self.y - other_point.y)
    
    def __neg__(self) -> 'Point':
        return Point(-self.x, -self.y)
    
    def to(self, other_point: 'Point') -> float:
        p = self - other_point
        return (p.x ** 2 + p.y ** 2) ** 0.5
    
    def __str__(self) -> str:
        return f'{self.__class__.__name__}({self.x}, {self.y})'

    def __repr__(self) -> str:
        return str(self)
    
    @property
    def is_center(self) -> bool:
        return self.x == 0 and self.y == 0
    
    @property
    def to_json(self) -> str:
        return json.dumps(self.__dict__)
    
    @classmethod
    def from_json(cls, s: str) -> 'Point':
        js = json.loads(s)
        return cls(js['x'], js['y'])