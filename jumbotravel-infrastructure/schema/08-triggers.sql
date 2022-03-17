CREATE DATABASE IF NOT EXISTS jumbotravel;
USE jumbotravel;

-- AirPlanes Trigger
DELIMITER //
CREATE TRIGGER airplane_update AFTER INSERT
ON master_airplanes FOR EACH ROW
BEGIN
    UPDATE master_airplanesmapping
    SET 
    updated_at = NOW()
    AND airplane_id = NEW.airplane_id
    WHERE
    1=1
    AND carrier = NEW.carrier
    AND flight_number = NEW.flight_number;
END
//

-- Agents Trigger
DELIMITER //
CREATE TRIGGER agent_update AFTER INSERT
ON master_agents FOR EACH ROW
BEGIN
    UPDATE master_agentmapping
    SET 
    updated_at = NOW()
    AND agent_id = NEW.agent_id
    WHERE
    1=1
    AND dni = NEW.dni;
END
//

-- Products Trigger
DELIMITER //
CREATE TRIGGER product_update AFTER INSERT
ON master_products FOR EACH ROW
BEGIN
    SET @product_id = (SELECT product_id FROM master_products WHERE product_code = NEW.product_code);
    UPDATE master_productmapping
    SET 
    updated_at = NOW()
    AND product_id = NEW.product_id
    WHERE
    1=1
    AND product_id = @product_id;
END
//

-- Stock Trigger
DELIMITER //
CREATE TRIGGER stock_update AFTER INSERT
ON airplane_stock FOR EACH ROW
BEGIN
    SET @stock_id = (SELECT stock_id FROM airplane_stock WHERE airplane_id = NEW.airplane_id AND product_id = NEW.product_id);
    UPDATE airplane_stockmapping
    SET 
    updated_at = NOW()
    AND stock_id = NEW.stock_id
    WHERE
    1=1
    AND stock_id = @stock_id;
END
//