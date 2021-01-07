--
-- SQL query that produces the invalid deliveries.
-- (Invalid deliveries are deliveries that a carrier cannot perform due to carrier bean constraints)
--

DO $$
DECLARE
  -- supplier id
  sid integer;
  -- carrier id
  cid integer;
  -- driver id
  did integer;
BEGIN
  FOR sid in select id from supplier 
  LOOP
    FOR cid in select carrier_id from carrier_bean_type where carrier_id not in
      (select carrier_id from carrier_bean_type where bean_type_id in (
         select bean_type_id from supplier_bean_type where supplier_id=sid
      )) GROUP BY carrier_id
      union
      select id from carrier where id not in (select carrier_id from carrier_bean_type)
    LOOP
      select id into did from driver where carrier_id=cid;
      if (did > 0) then
        insert into delivery (supplier_id, driver_id)
        values (sid, (select id from driver where carrier_id=cid));
      end if;
    END LOOP;
  END LOOP;
END; $$
