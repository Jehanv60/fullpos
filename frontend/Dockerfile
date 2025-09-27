# -- STAGE 1: Build the SvelteKit application --
FROM node:20-alpine AS build

# Set working directory
RUN mkdir /app
WORKDIR /app

# Copy package.json and install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the SvelteKit application for production
# Assuming you're using `@sveltejs/adapter-node` or another adapter that outputs a standalone app.
# If using adapter-static, adjust accordingly.
EXPOSE 3000

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]

# -- STAGE 2: Run the SvelteKit application --
# Use a lighter Node.js image for production
# FROM node:20-alpine

# # Set working directory
# WORKDIR /app
# COPY .env .

# # Copy the built application from the build stage
# COPY --from=build /app/build ./build
# COPY --from=build /app/package.json ./package.json

# # If using adapter-node, copy the entry point (e.g., index.js or handler.js)
# # Adjust this line based on your adapter's output.
# # For adapter-node, it typically creates a 'index.js' or 'handler.js' in the build directory.
# # Let's assume 'index.js' for simplicity, but check your adapter's output.
# COPY --from=build /app/node_modules ./node_modules

# # Expose the port SvelteKit listens on (default for adapter-node is 3000)
# EXPOSE 3000

# # Command to run the SvelteKit application
# # For adapter-node, you typically run the generated JS file.
# CMD ["node", "build"]