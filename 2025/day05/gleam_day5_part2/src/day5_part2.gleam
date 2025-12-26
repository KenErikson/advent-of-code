import gleam/io
import gleam/string
import simplifile

pub fn main() -> Nil {
  io.println("Hello from day5_part2!")

  let filepath = "../input_example.txt"

  case simplifile.read(filepath) {
    Ok(contents) -> {
      // 3. Split the content into a list of lines
      let lines = string.split(contents, on: "\n")

      // 4. Process the lines (e.g., print them)
      process_lines(lines)
    }
    Error(error) -> {
      io.println("Failed to read file: " <> simplifile.describe_error(error))
    }
  }
}

fn process_lines(lines: List(String)) {
  case lines {
    [] -> Nil
    // End of list
    [line, ..rest] -> {
      // Do something with the line
      io.println("Read line: " <> line)

      // Recursively process the rest
      process_lines(rest)
    }
  }
}
