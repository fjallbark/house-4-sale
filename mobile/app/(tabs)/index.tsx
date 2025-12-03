import { useEffect, useState } from "react";
import { FlatList, Text } from "react-native";
import fastighetsbyranData from "../../assets/fastighetsbyran-data.json";
import { ListingCard } from "@/components/ui/ListingCard";

type House = {
  title: string;
  price: string;
  rooms: string;
  url: string;
  squareMeters: string;
  date: string;
  address: string;
  image: string;
};

const isToday = (isoString: string) => {
  const d = new Date(isoString);
  const now = new Date();

  return (
    d.getFullYear() === now.getFullYear() &&
    d.getMonth() === now.getMonth() &&
    d.getDate() === now.getDate()
  );
};

export default function Index() {
  const [houses, setHouses] = useState<House[]>([]);

  const todayItems = houses.filter((item) => isToday(item.date));
  const otherItems = houses.filter((item) => !isToday(item.date));

  useEffect(() => {
    const sortByDate = (a: { date: string }, b: { date: string }) =>
      new Date(a.date).getTime() - new Date(b.date).getTime();
    // Currently I want to test against test data so I leave this like this for now
    /*fetch("http://localhost:8080/api/houses")
      .then((res) => res.json())
      .then((data) => setHouses(data.sort(sortByDate)))
      .catch(console.error);*/

    setHouses(fastighetsbyranData.sort(sortByDate));
  }, []);

  return (
    <>
      <Text style={{ fontSize: 20, fontWeight: "bold", marginBottom: 8 }}>
        Today
      </Text>
      <FlatList
        data={todayItems}
        keyExtractor={(_, index) => index.toString()}
        renderItem={({ item }) => <ListingCard item={item} />}
      />
      <Text style={{ fontSize: 20, fontWeight: "bold", marginBottom: 8 }}>
        Older
      </Text>
      <FlatList
        data={otherItems}
        keyExtractor={(_, index) => index.toString()}
        renderItem={({ item }) => <ListingCard item={item} />}
      />
    </>
  );
}
