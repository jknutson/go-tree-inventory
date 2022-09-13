CREATE OR REPLACE VIEW tree_inventory_view_v1 AS
	SELECT 
	  id,
	  type,
	  diameter_breast_height_inches,
	  diameter_breast_height_inches/39.37 AS diameter_breast_height_meters,
	  diameter_dripline_feet,
	  diameter_dripline_feet/3.281 AS diameter_dripline_meters,
	  ST_SetSRID(ST_MakePoint(SPLIT_PART(location, ',', 2)::float, SPLIT_PART(location, ',', 1)::float), 4269) AS geom
	FROM tree_inventory_v1
	WHERE location LIKE '%,%';
