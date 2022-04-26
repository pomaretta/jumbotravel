CREATE DATABASE IF NOT EXISTS jumbotravel;
USE jumbotravel;

CREATE TABLE IF NOT EXISTS invoices (

	invoice_id					INT NOT NULL AUTO_INCREMENT,

	agent_id					INT NOT NULL,
	agentmapping_id				INT NOT NULL,

	provider_id					INT NOT NULL,
	providermapping_id			INT NOT NULL,

	report_date 				DATE NULL,

    updated_at					TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	created_at					TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

	PRIMARY KEY (invoice_id),
	FOREIGN KEY (agent_id) REFERENCES master_agents(agent_id),
    FOREIGN KEY (agentmapping_id) REFERENCES master_agentmapping(agentmapping_id),
	FOREIGN KEY (provider_id) REFERENCES master_agents(agent_id),
    FOREIGN KEY (providermapping_id) REFERENCES master_agentmapping(agentmapping_id)
);

CREATE TABLE IF NOT EXISTS invoice_bookings (

	invoice_id					INT NOT NULL,
	bookingreferenceid			VARCHAR(255) NOT NULL,

	created_at					TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

	PRIMARY KEY (invoice_id, bookingreferenceid),
	FOREIGN KEY (invoice_id) REFERENCES invoices(invoice_id),
	FOREIGN KEY (bookingreferenceid) REFERENCES bookings(bookingreferenceid)
);