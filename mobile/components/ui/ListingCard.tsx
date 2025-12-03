import React from "react";
import { Text, Image, Pressable, Linking, StyleSheet } from "react-native";

// Typ för item
type Item = {
  url: string;
  image: string;
  title: string;
  price: string;
  rooms: string;
  squareMeters: string;
  address: string;
};

type Props = {
  item: Item;
};

export const ListingCard: React.FC<Props> = ({ item }) => {
  const handlePress = () => {
    Linking.openURL(item.url);
  };

  return (
    <Pressable onPress={handlePress} style={styles.container}>
      <Image source={{ uri: item.image }} style={styles.image} />
      <Text style={styles.title}>{item.title}</Text>
      <Text>{item.price}</Text>
      <Text>{item.rooms}</Text>
      <Text>{item.squareMeters}</Text>
      <Text>{item.address}</Text>
    </Pressable>
  );
};

const styles = StyleSheet.create({
  container: {
    padding: 15,
    backgroundColor: "#fff",
    borderRadius: 8,
    marginBottom: 12,
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 3,
    elevation: 2,
  },
  image: {
    width: "100%",
    height: 200,
    borderRadius: 8,
    marginBottom: 8,
  },
  title: {
    fontWeight: "bold",
    marginBottom: 4,
  },
});
