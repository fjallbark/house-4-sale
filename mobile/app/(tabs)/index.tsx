import { useEffect, useState } from "react";
import { View, Text, FlatList, Image } from "react-native";

type House = {
  title: string;
  price: string;
  address: string;
  image: string;
};

export default function Index() {
  const [houses, setHouses] = useState<House[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/api/houses")
      .then((res) => res.json())
      .then(setHouses)
      .catch(console.error);
  }, []);

  return (
    <FlatList
      data={houses}
      keyExtractor={(_, index) => index.toString()}
      renderItem={({ item }) => (
        <View style={{ padding: 15 }}>
          <Image
            source={{ uri: item.image }}
            style={{ width: "100%", height: 200 }}
          />
          <Text>{item.title}</Text>
          <Text>{item.price}</Text>
          <Text>{item.address}</Text>
        </View>
      )}
    />
  );
}
