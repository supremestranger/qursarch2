import React, { useState, useEffect } from "react";
import { useParams, Link } from "react-router-dom";

function MedicineDetail() {
  const { id } = useParams();
  const [drug, setDrug] = useState(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchDrugDetails = async () => {
      setLoading(true);
      try {
        const response = await fetch(`http://localhost:3001/v1/drugs/${id}`);
        const data = await response.json();
        setDrug(data);
      } catch (error) {
        console.error("Error fetching drug details:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchDrugDetails();
  }, [id]);

  if (loading) {
    return <p className="text-center text-gray-500">Loading...</p>;
  }

  if (!drug) {
    return <p className="text-center text-gray-500">Drug not found.</p>;
  }

  return (
    <div className="max-w-4xl mx-auto bg-white p-6 shadow-md rounded-md border border-gray-300">
      <Link to="/" className="text-blue-500 underline mb-4 inline-block">
        &larr; Back to List
      </Link>
      <h2 className="text-2xl font-bold mb-4">{drug.Name}</h2>
      <p className="text-gray-700 mb-4">{drug.Description}</p>
      <h3 className="text-xl font-semibold">Details</h3>
      <ul className="list-disc list-inside text-gray-600">
        <li>Требует рецепта: {drug.NeedsReceipt ? "Да" : "Нет"}</li>
        <li>Минимальный возраст: {drug.MinAge} years</li>
        <li>Вещества: {drug.Components?.join(", ")}</li>
        <li>Показания: {drug.Indications?.join(", ")}</li>
        <li>Вес: {drug.UnitWeight}</li>
        <li>Дозировка: {drug.Dosage}</li>
        <li>Эффекты: {drug.Effects}</li>
      </ul>
      <h3 className="text-xl font-semibold mt-6">Prices</h3>
      <ul className="mt-2">
        {drug.Prices?.sort((a, b) => a.Price - b.Price).map((price, index) => (
          <li
            key={index}
            className="flex justify-between border-b border-gray-200 pb-1 mb-1"
          >
            <span>{price.MedstoreName}</span>
            <span className="font-semibold text-green-500">{price.Price}₽</span>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default MedicineDetail;
