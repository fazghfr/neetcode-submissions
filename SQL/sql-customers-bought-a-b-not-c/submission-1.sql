-- Write your query below


-- WHERE IN A, B AND NOT IN C

select o.customer_id, o.customer_name
from customers o
where o.customer_id in (
    select customer_id
    from orders
    where orders.product_name = 'A'
) AND o.customer_id in (
    select customer_id
    from orders
    where orders.product_name = 'B'
) AND o.customer_id not in (
    select customer_id
    from orders
    where orders.product_name = 'C'
)
order by o.customer_name

