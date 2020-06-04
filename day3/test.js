// https://www.youtube.com/watch?v=A86COO8KC58

function segmentIntersect(p0, p1, p2, p3) {
    var A1 = p1.y - p0.y,
        B1 = p0.x - p1.x,
        C1 = A1 * p0.x + B1 * p0.y,
        A2 = p3.y - p2.y,
        B2 = p2.x - p3.x,
        C2 = A2 * p2.x + B2 * p2.y,
        denominator = A1 * B2 - A2 * B1;

    console.log('denominator', denominator);

    if (denominator == 0) {
        return null;
    }

    var intersectX = (B2 * C1 - B1 * C2) / denominator,
        intersectY = (A1 * C2 - A2 * C1) / denominator,
        rx0 = (intersectX - p0.x) / (p1.x - p0.x),
        ry0 = (intersectY - p0.y) / (p1.y - p0.y),
        rx1 = (intersectX - p2.x) / (p3.x - p2.x),
        ry1 = (intersectY - p2.y) / (p3.y - p2.y);

    console.log('rx0', rx0, 'ry0', ry0, 'rx1', rx1, 'ry1', ry1);
    // * Intersection line1={p{x:8,y:0} p{x:8,y:5}} line2={p{x:0,y:7} p{x:6,y:7}} at p{x:8,y:7}

    if (((rx0 >= 0 && rx0 <= 1) || (ry0 >= 0 && ry0 <= 1)) &&
        ((rx1 >= 0 && rx1 <= 1) || (ry1 >= 0 && ry1 <= 1))) {
        return {
            x: intersectX,
            y: intersectY
        };
    }
    else {
        return null;
    }
}


// result = segmentIntersect({ x: 2, y: 2 }, { x: 8, y: 2 }, { x: 3, y: 1 }, { x: 3, y: 5 })
result = segmentIntersect({ x: 8, y: 0 }, { x: 8, y: 5 }, { x: 0, y: 7 }, { x: 8, y: 7 })
console.log('res', result)

//x=3 y=3 ua=0 ub=0 line=l{start:p{x:3,y:5},end:p{x:3,y:2}} otherLine=l{start:p{x:6,y:3},end:p{x:2,y:3}}
result = segmentIntersect({ x: 3, y: 5 }, { x: 3, y: 2 }, { x: 6, y: 3 }, { x: 2, y: 3 })
console.log('res', result)