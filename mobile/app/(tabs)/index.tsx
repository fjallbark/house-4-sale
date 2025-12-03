import { useEffect, useState } from "react";
import { Text, SectionList } from "react-native";
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
  const noDate = houses.filter((item) => item.date === "");
  const otherItems = houses.filter(
    (item) => item.date !== "" && !isToday(item.date)
  );

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
      <SectionList
        sections={[
          {
            title: "Today",
            data: todayItems,
          },
          {
            title: "No Date",
            data: noDate,
          },
          {
            title: "Older",
            data: otherItems,
          },
        ].filter((section) => section.data.length > 0)}
        renderItem={({ item }) => <ListingCard item={item} />}
        renderSectionHeader={({ section }) => (
          <Text style={{ fontSize: 20, fontWeight: "bold", marginBottom: 8 }}>
            {section.title}
          </Text>
        )}
        keyExtractor={(item) => `basicListEntry-${item}`}
      />
    </>
  );
}
