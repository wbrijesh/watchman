export default function Table({
  columns,
  data,
}: {
  columns: string[];
  data: any[];
}) {
  return (
    <div>
      <div className="flow-root">
        <div className="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
          <div className="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
            <div className="overflow-hidden shadow ring-1 ring-slate-200 sm:-lg">
              <table className="min-w-full divide-y divide-slate-200">
                <thead className="bg-slate-100">
                  <tr>
                    {columns.map((column: string) => (
                      <th
                        key={column}
                        scope="col"
                        className="px-3 py-2 text-left text-sm font-semibold text-slate-500"
                      >
                        {column}
                      </th>
                    ))}
                  </tr>
                </thead>
                <tbody className="divide-y divide-slate-200 bg-white">
                  {data.map((row: any) => (
                    <tr key={row.id}>
                      {columns.map((column: string) => (
                        <td
                          key={column}
                          className="whitespace-nowrap py-2 px-3 text-sm text-slate-600"
                        >
                          {row[column.toLowerCase()]}
                        </td>
                      ))}
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
