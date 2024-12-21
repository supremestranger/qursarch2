import React, { useState, useEffect } from "react";

function App() {
  const [drugs, setDrugs] = useState([]);
  const [loading, setLoading] = useState(false);

  // Filter state
  const [filters, setFilters] = useState({
    name: "",
    needsReceipt: false,
    minAge: "",
    activeComponents: [],
    indications: [],
  });

  // Available options for Active Components and Indications
  const [activeComponentOptions, setActiveComponentOptions] = useState([]);
  const [indicationOptions, setIndicationOptions] = useState([]);

  // Fetch data with filters
  useEffect(() => {
    const fetchDrugs = async () => {
      setLoading(true);
      try {
        const queryParams = new URLSearchParams();
  
        if (filters.name) queryParams.append("name", filters.name);
        if (filters.needsReceipt) queryParams.append("needsReceipt", true);
        if (filters.minAge) queryParams.append("minAge", filters.minAge);
        
        // Send activeComponents as a comma-separated string
        if (filters.activeComponents.length > 0) {
          queryParams.append("activeComponents", filters.activeComponents.join(","));
        }
  
        if (filters.indications.length > 0) {
          queryParams.append("indications", filters.indications.join(","));
        }
  
        console.log(queryParams.toString())  // Check the generated query parameters
  
        const response = await fetch(`http://localhost:3001/v1/drugs?${queryParams.toString()}`);
        const data = await response.json();
        setDrugs(data);
      } catch (error) {
        console.error("Error fetching data:", error);
      } finally {
        setLoading(false);
      }
    };
  
    fetchDrugs();
  }, [filters]);

  // Fetch Active Components and Indications
  useEffect(() => {
    const fetchOptions = async () => {
      try {
        const activeComponentsResponse = await fetch("http://localhost:3001/v1/components");
        const componentsData = await activeComponentsResponse.json();
        setActiveComponentOptions(componentsData);

        const indicationsResponse = await fetch("http://localhost:3001/v1/indications");
        const indicationsData = await indicationsResponse.json();
        setIndicationOptions(indicationsData);
      } catch (error) {
        console.error("Error fetching options:", error);
      }
    };

    fetchOptions();
  }, []);

  // Update filters
  const handleFilterChange = (e) => {
    const { name, value, type, checked } = e.target;

    if (name === "activeComponents" || name === "indications") {
      setFilters((prevFilters) => ({
        ...prevFilters,
        [name]: checked
          ? [...prevFilters[name], value]
          : prevFilters[name].filter((item) => item !== value),
      }));
    } else if (type === "checkbox") {
      setFilters((prevFilters) => ({
        ...prevFilters,
        [name]: checked,
      }));
    } else {
      setFilters((prevFilters) => ({
        ...prevFilters,
        [name]: value,
      }));
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 text-gray-800">
      <header className="bg-blue-500 text-white py-4 px-6 shadow-md text-center">
        <h1 className="text-3xl font-bold">Medical Drug Market</h1>
      </header>
      <main className="flex p-6">
        {/* Side Panel for Filters */}
        <aside className="w-1/4 bg-white p-6 shadow-md border border-gray-300 rounded-md mr-4">
          <h2 className="text-2xl font-semibold mb-4">Filters</h2>
          <form className="space-y-4">
            {/* Name Filter */}
            <div>
              <label className="block text-gray-700 font-medium">Name:</label>
              <input
                type="text"
                name="name"
                value={filters.name}
                onChange={handleFilterChange}
                placeholder="Enter drug name"
                className="mt-1 p-2 border border-gray-300 rounded-md w-full"
              />
            </div>

            {/* Needs Receipt */}
            <div>
              <label className="inline-flex items-center">
                <input
                  type="checkbox"
                  name="needsReceipt"
                  checked={filters.needsReceipt}
                  onChange={handleFilterChange}
                  className="form-checkbox h-5 w-5 text-blue-600"
                />
                <span className="ml-2">Needs Receipt</span>
              </label>
            </div>

            {/* Minimum Age */}
            <div>
              <label className="block text-gray-700 font-medium">Minimum Age:</label>
              <input
                type="number"
                name="minAge"
                value={filters.minAge}
                onChange={handleFilterChange}
                className="mt-1 p-2 border border-gray-300 rounded-md w-full"
              />
            </div>

            {/* Active Components */}
            <div>
              <label className="block text-gray-700 font-medium mb-2">Active Components:</label>
              {activeComponentOptions.map((component) => (
                <label key={component} className="inline-flex items-center mr-4">
                  <input
                    type="checkbox"
                    name="activeComponents"
                    value={component}
                    checked={filters.activeComponents.includes(component)}
                    onChange={handleFilterChange}
                    className="form-checkbox h-5 w-5 text-blue-600"
                  />
                  <span className="ml-2">{component}</span>
                </label>
              ))}
            </div>

            {/* Indications */}
            <div>
              <label className="block text-gray-700 font-medium mb-2">Indications:</label>
              {indicationOptions.map((indication) => (
                <label key={indication} className="inline-flex items-center mr-4">
                  <input
                    type="checkbox"
                    name="indications"
                    value={indication}
                    checked={filters.indications.includes(indication)}
                    onChange={handleFilterChange}
                    className="form-checkbox h-5 w-5 text-blue-600"
                  />
                  <span className="ml-2">{indication}</span>
                </label>
              ))}
            </div>
          </form>
        </aside>

        {/* Drug List */}
        <section className="w-3/4 p-6 bg-white border border-gray-300 rounded-md shadow-md">
          {loading ? (
            <p className="text-center text-gray-500">Loading...</p>
          ) : drugs != null && drugs.length > 0 ? (
            <ul className="space-y-4">
              {drugs.map((drug) => (
              <li
                key={drug.Id}
                className="bg-white rounded-lg shadow-md p-4 border border-gray-300 hover:shadow-lg transition-shadow flex justify-between items-start"
              >
                <div>
                  <h2 className="text-xl font-semibold">{drug.Name}</h2>
                  <p className="text-gray-600 mt-2">{drug.Description}</p>
                  <p className={`mt-2 font-medium ${drug.NeedsReceipt ? "text-red-500" : "text-green-500"}`}>
                    {drug.NeedsReceipt ? "Требует рецепта" : "Не требует рецепта"}
                  </p>
                </div>
                <div className="text-right pr-4">
                <h3 className="text-lg font-semibold">Prices:</h3>
                <ul className="mt-2">
                  {drug.Prices?.sort((a, b) => a.Price - b.Price)
                    .slice(0, 3)
                    .map((price, index) => (
                      <li
                        key={index}
                        className="text-gray-800 text-sm flex justify-between border-b border-gray-200 pb-1 mb-1"
                      >
                        <span>{price.MedstoreName}</span>
                        <span className="font-semibold text-green-500">${price.Price}</span>
                      </li>
                    ))}
                </ul>
              </div>
              </li>
            ))}
            </ul>
          ) : (
            <p className="text-center text-gray-500">No drugs found.</p>
          )}
        </section>
      </main>
    </div>
  );
}

export default App;
