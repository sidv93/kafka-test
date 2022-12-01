import {
    createColumnHelper,
    flexRender,
    getCoreRowModel,
    useReactTable,
} from '@tanstack/react-table';
import { useEffect, useState } from 'react';

import './App.css';
import { SOCKET_SERVER_URL } from './constants/ApiRoutes';
import { sampleData } from './sampledata';


const columnHelper = createColumnHelper();
const columns = [
    columnHelper.accessor('MemTotal', {
        cell: info => info.getValue(),
        header: 'Total Memory',
    }),
    columnHelper.accessor('MemFree', {
        cell: info => info.getValue(),
        header: 'Memory Free',
    }),
    columnHelper.accessor('MemAvailable', {
        cell: info => info.getValue(),
        header: 'Memory available'
    }),
    columnHelper.accessor('SwapTotal', {
        cell: info => info.getValue(),
        header: 'Total Swap memory'
    }),
    columnHelper.accessor('SwapCached', {
        cell: info => info.getValue(),
        header: 'Cached swap memory'
    }),
    columnHelper.accessor('SwapFree', {
        cell: info => info.getValue(),
        header: 'Free swap memory'
    })
]

function App() {
    const [data, setData] = useState(sampleData);

    useEffect(() => {
        let socket = new WebSocket('ws://localhost:4000/socket');
        socket.onmessage = function (e) {
            if (!!e.data) {
                const data = JSON.parse(e.data);
                console.log('d', data);
                setData((d) => [...d, data]);
            }
        }
    }, []);

    const table = useReactTable({
        data,
        columns,
        getCoreRowModel: getCoreRowModel(),
    })

    return (
        <div className="App">
            <table>
                <thead>
                    {table.getHeaderGroups().map(headerGroup => (
                        <tr key={headerGroup.id}>
                            {headerGroup.headers.map(header => (
                                <th key={header.id}>
                                    {header.isPlaceholder
                                        ? null
                                        : flexRender(
                                            header.column.columnDef.header,
                                            header.getContext()
                                        )}
                                </th>
                            ))}
                        </tr>
                    ))}
                </thead>
                <tbody>
                    {table.getRowModel().rows.map(row => (
                        <tr key={row.id}>
                            {row.getVisibleCells().map(cell => (
                                <td key={cell.id}>
                                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                                </td>
                            ))}
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )
}

export default App
