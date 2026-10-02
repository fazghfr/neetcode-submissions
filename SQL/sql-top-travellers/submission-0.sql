-- Write your query below
SELECT u.name,
       COALESCE(SUM(r.distance), 0) AS travelled_distance
from rides r
right join users u on u.id = r.user_id
group by user_id, u.name
ORDER BY travelled_distance desc, name asc


