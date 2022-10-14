import React from "react";
import AnnonceCard from "../components/Card/AnnonceCard"
import { Flex, VStack } from "@chakra-ui/layout";
import TopBar2 from "../components/Bar/TopBarTwo"

export default function homeConnecte() {
    return (
        <>
            <TopBar2 />
            <Flex w="100%" justify="center" marginTop="20">
                <VStack w="100%" maxW="800px" >
                    <AnnonceCard />
                </VStack>
            </Flex>
        </>
    )
}