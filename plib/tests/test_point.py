import pytest
from plib import Point

@pytest.fixture
def points():
    return Point(0, 0), Point(2, 2)

@pytest.fixture
def point_and_json():
    return Point(1,1), '{"x": 1, "y": 1}'

class TestPoint:

    def test_creation(self):
        p = Point(1, 2)
        assert p.x == 1 and p.y == 2
        
        with pytest.raises(TypeError):
            Point(1.5, 1.5)

    def test_add(self, points):
        p1, p2 = points
        assert p2 + p1 == Point(2, 2)

    def test_iadd(self, points):
        p1, p2 = points
        p1 += p2
        assert p1 == Point(2, 2)

    def test_sub(self, points):
        p1, p2 = points
        assert p2 - p1 == Point(2, 2)
        assert p1 - p2 == -Point(2, 2)

    def test_isub(self, points):
        p1, p2 = points
        p2 -= p1
        assert p2 == Point(2, 2)

        p1, p2 = points
        p1 -= p2
        assert p1 == -Point(2, 2)
    
    def test_distance_to(self):
        p1 = Point(0, 0)
        p2 = Point(2, 0)
        assert p1.to(p2) == 2

    @pytest.mark.parametrize(
            'p1, p2, distance',
            [(Point(0, 0), Point(0, 10), 10),
             (Point(0, 0), Point(10, 0), 10),
             (Point(0, 0), Point(1, 1), 1.414)]
    )
    def test_distance_all_axis(self, p1, p2, distance):
        assert p1.to(p2) == pytest.approx(distance, 0.001)

    def test_converting_to_str(self):
        s = 'Point(0, 0)'
        p = Point(0, 0)
        assert str(p) == s

    def test_repr(self, points):
        reprs = '(Point(0, 0), Point(2, 2))'
        assert repr(points) == reprs

    def test_point_being_center(self, points):
        p1, p2 = points
        assert p1.is_center == True
        assert p2.is_center == False

    def test_converting_point_to_json(self, point_and_json):
        p, js = point_and_json
        assert p.to_json == js

    def test_getting_point_from_json(self, point_and_json):
        p, js = point_and_json
        assert p == Point.from_json(js)