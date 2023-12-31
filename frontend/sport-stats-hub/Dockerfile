# Start with a Node.js base image
FROM node:14

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json (or yarn.lock) to the image
COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json

# Install dependencies
RUN npm install

# Copy the rest of the application into the image
COPY . /app

# Expose the port the app runs on
EXPOSE 3000

# Start the application
CMD ["npm", "start"]
