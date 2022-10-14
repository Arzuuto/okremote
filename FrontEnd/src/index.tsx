import React from 'react';
import ReactDOM from 'react-dom';

import { ChakraProvider } from '@chakra-ui/react';

import Routess from './pages/routes';

import theme from './theme';
import './theme/index.css';

ReactDOM.render(
	<React.StrictMode>
		<ChakraProvider theme={theme} resetCSS>
			<Routess />
		</ChakraProvider>
	</React.StrictMode>,
	document.getElementById('root'),
);