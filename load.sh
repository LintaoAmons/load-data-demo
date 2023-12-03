#/bin/sh

for ((i = 0; i < 10; i++)); do
  PGPASSWORD=world123 psql -h localhost -p 5499 -d world-db -U world -c "\copy public.city from 'city-$i.csv' with (format csv);"
done
