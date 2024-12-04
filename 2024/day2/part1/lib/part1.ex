defmodule Part1 do
  @expected_record_count 1000

  def worker do
    {:ok, file} = File.open("./input.txt", [:utf8])

    for line <- IO.stream(file, :line) do
      pid = spawn(&check_line_messages/0)
      send(pid, {:line, line, self()})
    end

    File.close(file)

    values = check_return_messages()
    values |> Enum.count(&(&1 == true)) |> IO.inspect(label: "total")
  end

  defp check_return_messages(values \\ [], count \\ 1) do
    receive do
      {:result, result} ->
        if count == @expected_record_count do
          [result | values]
        else
          check_return_messages([result | values], count + 1)
        end

      {_} ->
        check_return_messages(values, count)
    end
  end

  defp check_line_messages() do
    receive do
      {:line, line, pid} ->
        line_as_ints = line |> String.trim() |> String.split() |> Enum.map(&String.to_integer(&1))
        send(pid, {:result, safe_line?(line_as_ints)})

      {_} ->
        check_line_messages()
    end
  end

  defp safe_line?(line, current_pos \\ nil)

  defp safe_line?(line, nil) do
    [num1 | num2] = line |> get_nums
    direction = check_direction(num1, num2)

    safe_line?(line, direction)
  end

  defp safe_line?(_, :none), do: false

  defp safe_line?(line, current_pos) do
    [num1 | num2] = line |> get_nums

    if num2 == [nil] do
      # reached the end with no issues
      true
    else
      if check_diff(num1, num2) && check_direction(num1, num2) == current_pos do
        # diff is good, and we didnt change directions
        [_ | tail] = line
        safe_line?(tail, current_pos)
      else
        false
      end
    end
  end

  defp check_diff(num1, num2) do
    diff =
      if num1 > num2 do
        num1 - num2
      else
        num2 - num1
      end

    3 >= diff && diff > 0
  end

  defp check_direction(num1, num2) do
    if num1 < num2 do
      :positive
    else
      if num1 > num2 do
        :negative
      else
        :none
      end
    end
  end

  defp get_nums([num1 | []]), do: [num1, nil]

  defp get_nums(line) do
    [num1 | tail] = line
    [num2 | _] = tail
    [num1 | num2]
  end
end
