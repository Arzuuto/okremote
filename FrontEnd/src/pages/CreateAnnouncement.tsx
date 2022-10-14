import React from "react";
import {
  HStack,
  Select,
  Image,
  Button,
  Flex,
  VStack,
  Box,
  Stack,
  FormControl,
  FormLabel,
  FormHelperText,
  Input,
  useToast,
} from "@chakra-ui/react";
import { useState } from "react";
import TopBar from "../components/Bar/TopBarTwo";
import { useNavigate } from "react-router-dom";

export default function CreateAnnouncement() {
  const toast = useToast();
  const navigate = useNavigate();
  const [post_name, setPost] = useState("");
  const [enterprise, setEnterprise] = useState("");
  const [sector, setSector] = useState("");
  const [contrat, setContrat] = useState("");
  const [location, setLocation] = useState("");
  const [remuneration, setRemuneration] = useState("");
  const [date, setDate] = useState("");
  const [image, setImage] = useState("https://logodix.com/logo/1254764.jpg");

  const handleClick = async () => {
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        post_name,
        enterprise,
        sector,
        contrat,
        location,
        remuneration,
        date,
        image
      }),
    };
    console.log(
      JSON.stringify({
        post_name,
        enterprise,
        sector,
        contrat,
        location,
        remuneration,
        date,
        image,
      }),
    );

    fetch("http://localhost:8080/CreateAnnonce", requestOptions)
      .then((data) => {
        console.log(data);
        navigate("/homeConnecte");
        toast({
          title: "Annonce créée!",
          status: "success",
          duration: 5000,
          isClosable: true,
        });
      })
      .catch((error) => {
        console.log(error);
        toast({
          title: "error",
          status: "error",
          duration: 5000,
          isClosable: true,
        });
      });
  };
  const btn =
    post_name === "" ||
    enterprise === "" ||
    sector === "" ||
    contrat === "" ||
    location === "" ||
    remuneration === "" ||
    date === "" ? (
      <Button
        disabled
        onClick={handleClick}
        colorScheme="green"
        type="submit"
        w="xl"
      >
        Soumettre
      </Button>
    ) : (
      <Button onClick={handleClick} colorScheme="green" type="submit" w="xl">
        Soumettre
      </Button>
    );

  return (
    <>
      <TopBar />
      <Flex marginTop="25px" w="70%" align="center" justify="center">
        <VStack align="start" w="70%" maxW="400px"></VStack>
        <Box
          bg="whiteAlpha.900"
          w="100%"
          p={50}
          color="black"
          borderWidth={5}
          borderRadius={50}
          px="10"
        >
          <Stack align="end">
            <HStack>
              <FormControl>
                <FormLabel
                  marginTop="0px"
                  w="50%"
                  placeholder="Entrez le nom de l'entreprise"
                  color="black"
                >
                  Nom du poste
                </FormLabel>
                <Input
                  onChange={(e) => setEnterprise(e.target.value)}
                  id="job name"
                  type="enterprise_name"
                  placeholder="Entrez le poste"
                  color="black"
                />
                <FormHelperText>Quelle est nom de l'entreprise</FormHelperText>
                <FormLabel marginTop="0px" w="50%" color="black">
                  Nom de l'entreprise
                </FormLabel>
                <Input
                  onChange={(e) => setPost(e.target.value)}
                  id="job name"
                  type="post_name"
                  placeholder="Entrez le nom de l'entreprise"
                  color="black"
                />
                <FormHelperText>Quelle est l'intulé du poste</FormHelperText>
                <FormLabel marginTop="10px" w="50%">
                  Secteur d'activité
                </FormLabel>
                <Input
                  onChange={(e) => setSector(e.target.value)}
                  placeholder="Entrez le nom du secteur"
                  id="job name"
                  type="sector"
                />
                <FormHelperText>
                  Informatique ? Restauration ? Choisissez !
                </FormHelperText>
                <Select
                  onChange={(e) => setContrat(e.target.value)}
                  marginTop="20px"
                  placeholder="Selectionez le contrat"
                  color="black"
                >
                  Type de contrat
                  <option value="cdd" color="black">
                    CDD
                  </option>
                  <option value="cdi">CDI</option>
                  <option value="stage">Stage</option>
                  <option value="alternance">Alternance</option>
                </Select>
                <FormLabel marginTop="10px" w="50%">
                  Début de contrat
                </FormLabel>
                <Input
                  onChange={(e) => setDate(e.target.value)}
                  id="job name"
                  type="location"
                  placeholder="Entrez la rémunération du salarié"
                  color="black"
                />
                <FormHelperText>
                  A partir de quand le contrat commence
                </FormHelperText>
                <FormLabel marginTop="10px" w="50%">
                  Rémunération
                </FormLabel>
                <Input
                  onChange={(e) => setRemuneration(e.target.value)}
                  id="job name"
                  type="location"
                  placeholder="Entrez la rémunération du salarié"
                  color="black"
                />
                <FormHelperText>
                  Valeur en euro par mois
                </FormHelperText>
                <FormLabel marginTop="10px" w="50%">
                  Localisation
                </FormLabel>
                <Input
                  onChange={(e) => setLocation(e.target.value)}
                  id="job name"
                  type="location"
                  placeholder="Entrez la localisation"
                  color="black"
                />
                <FormHelperText>Où se situe le poste a occupé</FormHelperText>
                <HStack px="240px" py="15">
                  <Image
                    alt="Image de l'entreprise"
                    boxSize="150px"
                    objectFit="cover"
                    borderRadius={100}
                    align="center"
                    src={image}
                  />
                </HStack>
                {btn}
              </FormControl>
            </HStack>
          </Stack>
        </Box>
      </Flex>
    </>
  );
}
