import React from "react";
import TopBar from "../components/Bar/TopBar"
import AnnonceCard from "../components/Card/AnnonceCard"
import { Flex, VStack } from "@chakra-ui/layout";

export default function Home() {
  return (
    <>
      <TopBar />
      <Flex w="100%" justify="center" marginTop="20">
        <VStack w="100%" maxW="800px" >
          <AnnonceCard />
        </VStack>
      </Flex>
    </>
  )
}