const firstNames = [
  'Aiden',
  'Sophia',
  'Jackson',
  'Olivia',
  'Liam',
  'Emma',
  'Lucas',
  'Ava',
  'Noah',
  'Isabella',
  'Ethan',
  'Mia',
  'Mason',
  'Amelia',
  'James',
  'Harper',
];

const lastNames = [
  'Smith',
  'Johnson',
  'Williams',
  'Jones',
  'Brown',
  'Davis',
  'Miller',
  'Wilson',
  'Moore',
  'Taylor',
  'Anderson',
  'Thomas',
  'Jackson',
  'White',
];

function getRandomElement<T>(arr: T[]): T {
  const randomIndex = Math.floor(Math.random() * arr.length);
  return arr[randomIndex];
}

export default function generateRandomName(): string {
  const firstName = getRandomElement(firstNames);
  const lastName = getRandomElement(lastNames);

  return `${firstName} ${lastName}`;
}
