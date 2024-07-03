import { BarChart, Card, Divider, Switch } from "@tremor/react";
import { useState } from "react";

function valueFormatter(number: any) {
  return new Intl.NumberFormat("en-US", {
    maximumFractionDigits: 0,
    notation: "compact",
    compactDisplay: "short",
    style: "currency",
    currency: "INR",
  }).format(number);
}

export default function BarChartExample() {
  const [showComparison, setShowComparison] = useState(false);
  return (
    <>
      {/* <Card className="ax-w-2xl"> */}
      <BarChart
        data={data}
        index="date"
        categories={["value"]}
        colors={["slate"]}
        valueFormatter={valueFormatter}
        showYAxis={false}
        className="my-4 h-40"
        showLegend={false}
        showGridLines={false}
        showAnimation={true}
        animationDuration={100}
        tickGap={30}
        barCategoryGap={3}
      />
      {/* </Card> */}
    </>
  );
}

const data = [
  { date: "May 23", value: 56000 },
  { date: "Jun 23", value: 30000 },
  { date: "Jul 23", value: 85390 },
  { date: "Aug 23", value: 80100 },
  { date: "Sep 23", value: 75090 },
  { date: "Oct 23", value: 71080 },
  { date: "Nov 23", value: 61210 },
  { date: "Dec 23", value: 60143 },
  { date: "Jan 24", value: 10000 },
  { date: "Feb 24", value: 10000 },
  { date: "Mar 24", value: 10000 },
  { date: "Apr 24", value: 10000 },
  { date: "May 24", value: 10000 },
  { date: "Jun 24", value: 10000 },
  { date: "Jul 24", value: 10000 },
  { date: "Aug 24", value: 10000 },
  { date: "Sep 24", value: 10000 },
  { date: "Oct 24", value: 10000 },
  { date: "Nov 24", value: 10000 },
  { date: "Dec 24", value: 10000 },
  { date: "Jan 25", value: 10000 },
  { date: "Feb 25", value: 10000 },
  { date: "Mar 25", value: 10000 },
  { date: "Apr 25", value: 10000 },
  { date: "May 25", value: 10000 },
  { date: "Jun 25", value: 10000 },
  { date: "Jul 25", value: 10000 },
  { date: "Aug 25", value: 10000 },
  { date: "Sep 25", value: 10000 },
  { date: "Oct 25", value: 10000 },
  { date: "Nov 25", value: 10000 },
  { date: "Dec 25", value: 10000 },
  { date: "Jan 26", value: 10000 },
  { date: "Feb 26", value: 10000 },
  { date: "Mar 26", value: 10000 },
  { date: "Apr 26", value: 10000 },
  { date: "May 26", value: 10000 },
  { date: "Jun 26", value: 10000 },
  { date: "Jul 26", value: 10000 },
  { date: "Aug 26", value: 10000 },
  { date: "Sep 26", value: 10000 },
  { date: "Oct 26", value: 10000 },
  { date: "Nov 26", value: 10000 },
  { date: "Dec 26", value: 10000 },
  { date: "Jan 27", value: 10000 },
  { date: "Feb 27", value: 10000 },
  { date: "Mar 27", value: 10000 },
  { date: "Apr 27", value: 10000 },
  { date: "May 27", value: 10000 },
  { date: "Jun 27", value: 10000 },
  { date: "Jul 27", value: 43523 },
  { date: "Aug 27", value: 10000 },
  { date: "Sep 27", value: 10000 },
  { date: "Oct 27", value: 10000 },
  { date: "Nov 27", value: 10000 },
  { date: "Dec 27", value: 10000 },
  { date: "Jan 28", value: 10000 },
  { date: "Feb 28", value: 10000 },
  { date: "Mar 28", value: 10000 },
  { date: "Apr 28", value: 10000 },
  { date: "May 28", value: 10000 },
  { date: "Jun 28", value: 10000 },
  { date: "Jul 28", value: 10000 },
  { date: "Aug 28", value: 10000 },
  { date: "Sep 28", value: 10000 },
  { date: "Oct 28", value: 10000 },
  { date: "Nov 28", value: 10000 },
  { date: "Dec 28", value: 10000 },
  { date: "Jan 29", value: 10000 },
  { date: "Feb 29", value: 10000 },
  { date: "Mar 29", value: 10000 },
  { date: "Apr 29", value: 10000 },
  { date: "May 29", value: 10000 },
  { date: "Jun 29", value: 10000 },
  { date: "Jul 29", value: 10000 },
  { date: "Aug 29", value: 10000 },
  { date: "Sep 29", value: 10000 },
  { date: "Oct 29", value: 10000 },
  { date: "Nov 29", value: 10000 },
  { date: "Dec 29", value: 10000 },
  { date: "Jan 30", value: 10000 },
  { date: "Feb 30", value: 10000 },
  { date: "Mar 30", value: 10000 },
  { date: "Apr 30", value: 10000 },
];
