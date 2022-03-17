CREATE DATABASE IF NOT EXISTS jumbotravel;
USE jumbotravel;

-- AirPlanes Trigger
DELIMITER //
CREATE TRIGGER airplane_update AFTER INSERT
ON master_airplanes FOR EACH ROW
mp: BEGIN
    
    -- Select AirPlane ID
    SELECT
    	mam.airplane_id INTO @airplane_id 
    FROM
        master_airplanesmapping mam
    LEFT JOIN master_airplanes ma
        ON
        ma.airplane_id = mam.airplane_id
    WHERE
        1 = 1
        AND ma.carrier = NEW.carrier
        AND ma.flight_number = NEW.flight_number
    ORDER BY
        ma.created_at DESC
    LIMIT 1;

    
    -- Create Airplane Mapping if not exists
    IF @airplane_id IS NULL THEN
        INSERT INTO master_airplanesmapping (airplane_id) VALUES (NEW.airplane_id);
        LEAVE mp;
    END IF;
    
    -- Update Airplane Mapping
    UPDATE master_airplanesmapping
    SET 
        airplane_id = NEW.airplane_id
    WHERE
        1=1
        AND airplane_id = @airplane_id;
END
//

-- Agents Trigger
DELIMITER //
CREATE TRIGGER agent_update AFTER INSERT
ON master_agents FOR EACH ROW
mp: BEGIN

    -- Select Agent ID
    SELECT
    	mam.agent_id INTO @agent_id
    FROM
        master_agentmapping mam
    LEFT JOIN master_agents ma
        ON
        ma.agent_id = mam.agent_id
    WHERE
        1=1
        AND ma.dni = NEW.dni
    ORDER BY
        ma.created_at DESC
    LIMIT 1;

    -- Create Agent Mapping if not exists
    IF @agent_id IS NULL THEN
        INSERT INTO master_agentmapping (agent_id) 
        VALUES (NEW.agent_id);
        LEAVE mp;
    END IF;
    
    -- Update Agent Mapping
    UPDATE master_agentmapping
    SET 
        agent_id = NEW.agent_id
    WHERE
        1=1
        AND agent_id = @agent_id;
END
//

-- Products Trigger
DELIMITER //
CREATE TRIGGER product_update AFTER INSERT
ON master_products FOR EACH ROW
mp: BEGIN
    
    -- Select Product ID
    SELECT
    	mp.product_id INTO @product_id
    FROM
        master_productmapping mpm
    LEFT JOIN master_products mp
        ON
        mp.product_id = mpm.product_id
    WHERE
        1=1
        AND mp.product_code = NEW.product_code
    ORDER BY
        mp.created_at DESC
    LIMIT 1;

    -- Create Product Mapping if not exists
    IF @product_id IS NULL THEN
        INSERT INTO master_productmapping (product_id) VALUES (NEW.product_id);
        LEAVE mp;
    END IF;
    
    -- Update Product Mapping
    UPDATE master_productmapping
    SET 
        product_id = NEW.product_id
    WHERE
        1=1
        AND product_id = @product_id;
END
//

-- Stock Trigger
DELIMITER //
CREATE TRIGGER stock_update AFTER INSERT
ON airplane_stock FOR EACH ROW
mp: BEGIN

    -- Select Stock ID
    SELECT
    	ast.stock_id INTO @stock_id
    FROM
        airplane_stockmapping asm
    LEFT JOIN airplane_stock ast
        ON
        ast.stock_id = asm.stock_id
    WHERE
        1=1
        AND ast.airplane_id = NEW.airplane_id
        AND ast.product_id = NEW.product_id
    ORDER BY
        ast.created_at DESC
    LIMIT 1;

    -- Create Stock Mapping if not exists
    IF @stock_id IS NULL THEN
        INSERT INTO airplane_stockmapping (stock_id) VALUES (NEW.stock_id);
        LEAVE mp;
    END IF;
    
    -- Update Stock Mapping
    UPDATE airplane_stockmapping
    SET 
        stock_id = NEW.stock_id
    WHERE
        1=1
        AND stock_id = @stock_id;
END
//