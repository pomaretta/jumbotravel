import { Stat, Composite } from "../domain/agent_stats";

class StatCollection {

    constructor({
        stats = [],
    }) {
        this.stats = stats;
    }

    getValues(valueName = "value") {
        return this.stats.map(stat => {
            let statData = {
                name: stat.name,
            }
            statData[valueName] = stat.value;
            return statData;
        });
    }

    getDays(
        days = 31,
        valueName = "Flights"
    ) {
        let data = [];
        let utcNow = new Date();
        
        // Iterate in the range of days
        for (let i = days; i >= 0; i--) {
            let currentDate = new Date(utcNow.getTime() - (i * 24 * 60 * 60 * 1000));
            // Check if there's a composite for the current day
            let stat = this.stats.find(stat => {
                // Parse the date of the composite name
                let statDate = new Date(stat.name);
                // Check if the composite date is the same as the current day
                return statDate.getDate() === currentDate.getDate() &&
                    statDate.getMonth() === currentDate.getMonth() &&
                    statDate.getFullYear() === currentDate.getFullYear();
            });

            let dateYear = currentDate.getFullYear();
            let dateMonth = ("0" + (currentDate.getMonth() + 1)).slice(-2)
            let dateDay = ("0" + (currentDate.getDate())).slice(-2)

            let statData = {
                name: `${dateYear}-${dateMonth}-${dateDay}`,
            }
            statData[valueName] = stat ? stat.value : 0;
            data.push(statData);
        }

        return data;
    }

    static parse(data) {
        return new StatCollection({
            stats: data.map(stat => new Stat(stat))
        });
    }

}

class CompositeCollection {

    constructor({
        composites = [],
    }) {
        this.composites = composites;
    }

    getDays(
        days = 30
    ) {

        let data = [];
        let utcNow = new Date();
        
        // Iterate in the range of days
        for (let i = days; i >= 0; i--) {

            let currentDate = new Date(utcNow.getTime() - (i * 24 * 60 * 60 * 1000));

            // Check if there's a composite for the current day
            let composite = this.composites.find(composite => {
                
                // Parse the date of the composite name
                let compositeDate = new Date(composite.name);

                // Check if the composite date is the same as the current day
                return compositeDate.getDate() === currentDate.getDate() &&
                    compositeDate.getMonth() === currentDate.getMonth() &&
                    compositeDate.getFullYear() === currentDate.getFullYear();
            });

            let dateYear = currentDate.getFullYear();
            let dateMonth = ("0" + (currentDate.getMonth() + 1)).slice(-2)
            let dateDay = ("0" + (currentDate.getDate())).slice(-2)

            data.push({
                "name": `${dateYear}-${dateMonth}-${dateDay}`,
                "Flights": composite ? composite.flights : 0,
                "Bookings": composite ? composite.bookings : 0,
                "Total": composite && composite.total ? composite.total : 0
            });
        }

        return data;
    }

    static parse(data) {
        return new CompositeCollection({
            composites: data.map(composite => new Composite(composite))
        });
    }

}

export { StatCollection, CompositeCollection };