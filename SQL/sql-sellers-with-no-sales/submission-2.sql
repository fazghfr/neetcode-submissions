

select seller_name
from seller
where seller_id not in (
    select seller_id
    from orders o
    where extract(year from sale_date) = 2020
)
order by seller_name asc


