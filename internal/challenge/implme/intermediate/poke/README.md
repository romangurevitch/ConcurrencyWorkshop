# Pokémon Challenge: Intermediate Asynchronous Method Implementation

## Disclaimer

**Intermediate Challenge**: This challenge is designed for developers who have a foundational understanding of Go's
concurrency model and are looking to apply these concepts in a practical setting.

Participants are expected to have prior experience with Go's basic features before attempting this challenge.  
This is an opportunity to apply your knowledge of goroutines and channels to improve application performance.

## Overview

The Pokémon Challenge invites participants to enhance the responsiveness of a Pokémon application by implementing an
intermediate-level asynchronous method, `OnChangeNonBlocking`. This method is essential for fetching data from the
PokéAPI without causing UI freezes or delays.

<img src="https://raw.githubusercontent.com/PokeAPI/media/master/logo/pokeapi_256.png" alt="drawing" height="150"/>

## Challenge Description

The existing Pokémon application performs data fetching synchronously, which may cause the UI to become unresponsive
during long-running operations. Your challenge is to convert this synchronous process into an asynchronous one using
Go's concurrency features.

Your task is to implement the `OnChangeNonBlocking` method to initiate data fetching in a way that allows the UI to
remain interactive.

## Provided Materials

- **Synchronous Pokémon client application**: A starting point with a blocking fetch mechanism.
- **Access to PokéAPI**: The application utilizes this public API to retrieve Pokémon data.
- **Guiding Documentation**: Comments and structure in the existing code to help you understand the application's
  workflow.
- **Initial Unit Tests**: To ensure that the synchronous implementation is currently functional.

## Goals

- Implement the `OnChangeNonBlocking` method in the [app package](app/app.go) to fetch Pokémon data asynchronously.
- Ensure that the UI does not block or freeze during data retrieval.
- Maintain the overall functionality and integrity of the application.

## Getting Started

1. Run the following command in your terminal (in the current directory) to start the application:

   ```shell
   go run main.go
   ```

2. Examine the existing synchronous data retrieval method in the application to understand how it currently operates.
3. Design a solution that employs on of the concurrent patterns to handle data fetching without blocking the UI.
4. Implement the `OnChangeNonBlocking` method according to your design.
5. Test your implementation to confirm that the UI stays responsive and that data fetching works as expected.
